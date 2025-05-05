package presistence

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/HasanNugroho/go-broilerplate-ddd/internal/domain/account"
	"github.com/HasanNugroho/go-broilerplate-ddd/internal/sharekernel/errs"
	"github.com/HasanNugroho/go-broilerplate-ddd/internal/sharekernel/identity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

/**
 * GetALl retrieves a user by filter.
 * @param ctx context.Context
 * @param limit int
 * @param page int
 * @param sort string
 * @param search string
 * @return (*UserResponse, error)
 */
func (u *UserRepository) GetAll(ctx context.Context, search string, limit int, page int, sort string) (result []*account.User, totalItems int64, err error) {
	offset := (page - 1) * limit

	sortField := "updated_at"
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	orderBy := fmt.Sprintf("%s %s", sortField, sort)

	query := u.db.WithContext(ctx).Model(&result).Order(orderBy).Limit(limit).Offset(offset)

	if search != "" {
		searchPattern := fmt.Sprintf("%%%s%%", search)
		query = query.Where("name ILIKE ? OR email ILIKE ?", searchPattern, searchPattern)
	}

	if err = query.Find(&result).Error; err != nil {
		return nil, 0, fmt.Errorf("users not found: %w", errs.ErrNotFound)
	}

	if err := query.Model(result).Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

/**
 * GetByID retrieves a user by their ID.
 * @param ctx context.Context
 * @param id identity.ID
 * @return (*User, error)
 */
func (u *UserRepository) GetByID(ctx context.Context, id identity.ID) (result *account.User, err error) {
	if err := u.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID '%s' not found: %w", id, errs.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to query user: %w", err)
	}
	return result, nil
}

/**
 * GetByEmail retrieves a user by their email.
 * @param ctx context.Context
 * @param email string
 * @return (*User, error)
 */
func (u *UserRepository) GetByEmail(ctx context.Context, email string) (result *account.User, err error) {
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with email '%s' not found: %w", email, errs.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to query user: %w", err)
	}
	return result, nil
}

/**
 * GetByUsername retrieves a user by their username.
 * @param ctx context.Context
 * @param username string
 * @return (*User, error)
 */
func (u *UserRepository) GetByUsername(ctx context.Context, username string) (result *account.User, err error) {
	if err := u.db.WithContext(ctx).Where("username = ?", username).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with username '%s' not found: %w", username, errs.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to query user: %w", err)
	}
	return result, nil
}

/**
 * Create creates a new user.
 * @param ctx context.Context
 * @param user *User
 * @return error
 */
func (u *UserRepository) Create(ctx context.Context, user *account.User) (err error) {
	result := u.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("user with email '%s' already exists: %w", user.Email, errs.ErrConflict)
		}
		return fmt.Errorf("failed to create user: %w", result.Error)
	}
	return nil
}

/**
 * Update updates a user.
 * @param ctx context.Context
 * @param user *User
 * @return (*User, error)
 */
func (u *UserRepository) Update(ctx context.Context, user *account.User) (err error) {
	result := u.db.WithContext(ctx).Save(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with ID '%s' not found: %w", user.ID, errs.ErrNotFound)
		}
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			if strings.Contains(result.Error.Error(), "username") {
				return fmt.Errorf("username already exists: %w", errs.ErrConflict)
			}
			if strings.Contains(result.Error.Error(), "email") {
				return fmt.Errorf("email already exists: %w", errs.ErrConflict)
			}
			return fmt.Errorf("duplicated key error: %w", errs.ErrConflict)
		}
		return fmt.Errorf("failed to update user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user with ID '%s' not found: %w", user.ID, errs.ErrNotFound)
	}
	return nil
}

/**
 * Delete deletes a user.
 * @param ctx context.Context
 * @param id identity.ID
 * @return error
 */
func (u *UserRepository) Delete(ctx context.Context, id identity.ID) (err error) {
	result := u.db.WithContext(ctx).Where("id", id).Delete(&account.User{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with ID '%s' not found: %w", id, errs.ErrNotFound)
		}
		return fmt.Errorf("failed to delete user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user with ID '%s' not found: %w", id, errs.ErrNotFound)
	}
	return nil
}
