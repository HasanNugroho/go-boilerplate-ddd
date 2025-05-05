package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/HasanNugroho/go-broilerplate-ddd/internal/domain/account"
	"github.com/HasanNugroho/go-broilerplate-ddd/internal/domain/account/interfaces"
	"github.com/HasanNugroho/go-broilerplate-ddd/internal/sharekernel/errs"
	"github.com/HasanNugroho/go-broilerplate-ddd/internal/sharekernel/identity"
	"github.com/google/uuid"
)

type UserService struct {
	repo interfaces.IUserRepository
}

func NewUserService(repo interfaces.IUserRepository) *UserService {
	return &UserService{repo: repo}
}

/**
 * GetALl retrieves a user by filter.
 * @param ctx context.Context
 * @param limit int
 * @param page int
 * @param sort string
 * @param search string
 * @return (*account.UserResponse, error)
 */
func (u *UserService) GetAll(ctx context.Context, search string, limit int, page int, sort string) (result []*account.UserResponse, totalItems int64, err error) {
	users, totalItem, err := u.repo.GetAll(ctx, search, limit, page, sort)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, 0, errs.ErrNotFound
		}
		return nil, 0, fmt.Errorf("failed to get users: %w", err)
	}

	if totalItem == 0 {
		return result, 0, errs.ErrNotFound
	}

	result = make([]*account.UserResponse, 0, len(users))

	for _, user := range users {
		userResponse := user.ToUserResponse()
		result = append(result, userResponse)
	}

	return result, totalItem, nil
}

/**
 * GetByID retrieves a user by their ID.
 * @param ctx context.Context
 * @param id identity.ID
 * @return (*account.UserResponse, error)
 */
func (u *UserService) GetByID(ctx context.Context, id identity.ID) (result *account.UserResponse, err error) {
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	result = user.ToUserResponse()

	return result, nil
}

/**
 * GetByEmail retrieves a user by their email.
 * @param ctx context.Context
 * @param email string
 * @return (*account.UserResponse, error)
 */
func (u *UserService) GetByEmail(ctx context.Context, email string) (result *account.UserResponse, err error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	result = user.ToUserResponse()

	return result, nil
}

/**
 * GetByUsername retrieves a user by their username.
 * @param ctx context.Context
 * @param username string
 * @return (*account.UserResponse, error)
 */
func (u *UserService) GetByUsername(ctx context.Context, username string) (result *account.UserResponse, err error) {
	user, err := u.repo.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	result = user.ToUserResponse()

	return result, nil
}

/**
 * Create creates a new user.
 * @param ctx context.Context
 * @param user *account.CreateUserRequest
 */
func (u *UserService) Create(ctx context.Context, user *account.CreateUserRequest) (err error) {
	newUser := &account.User{
		Name:     user.Name,
		Username: user.Username,
		Fullname: user.Fullname,
		Email:    user.Email,
	}
	newUser.EncryptPassword(user.Password)

	if err = u.repo.Create(ctx, newUser); err != nil {
		if errors.Is(err, errs.ErrConflict) {
			return errs.ErrConflict
		}
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

/**
 * Update updates a user.
 * @param ctx context.Context
 * @param id identity.ID
 * @param user *UpdateUserRequest
 * @return error
 */
func (u *UserService) Update(ctx context.Context, id identity.ID, payload *account.UpdateUserRequest) (err error) {
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrNotFound
		}
		return fmt.Errorf("failed to get user by id: %w", err)
	}

	if payload.Name != "" {
		user.Name = payload.Name
	}

	if payload.Fullname != "" {
		user.Fullname = payload.Fullname
	}

	if payload.Username != "" {
		user.Username = payload.Username
	}

	if payload.Email != "" {
		user.Email = payload.Email
	}

	if payload.Password != "" {
		if err = user.EncryptPassword(payload.Password); err != nil {
			return fmt.Errorf("failed to encrypt password: %w", errs.ErrInternal)
		}
	}

	if uuid.UUID(payload.Role) != uuid.Nil {
		user.Role = payload.Role
	}

	if err = u.repo.Update(ctx, user); err != nil {
		if errors.Is(err, errs.ErrConflict) {
			return errs.ErrConflict
		}
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil

}

/**
 * Delete deletes a user.
 * @param ctx context.Context
 * @param id identity.ID
 * @return error
 */
func (u *UserService) Delete(ctx context.Context, id identity.ID) (err error) {
	_, err = u.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrNotFound
		}
		return fmt.Errorf("failed to get user by id: %w", err)
	}

	if err = u.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrNotFound
		}
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
