package interfaces

import (
	"context"

	"github.com/HasanNugroho/go-broilerplate-ddd/internal/domain/account"
	"github.com/HasanNugroho/go-broilerplate-ddd/internal/sharekernel/identity"
)

type IUserService interface {
	/**
	 * GetALl retrieves a user by filter.
	 * @param ctx context.Context
	 * @param limit int
	 * @param page int
	 * @param sort string
	 * @param search string
	 * @return (*UserResponse, error)
	 */
	GetAll(ctx context.Context, search string, limit int, page int, sort string) (result []*account.UserResponse, totalItems int64, err error)

	/**
	 * GetByID retrieves a user by their ID.
	 * @param ctx context.Context
	 * @param id identity.ID
	 * @return (*UserResponse, error)
	 */
	GetByID(ctx context.Context, id identity.ID) (result *account.UserResponse, err error)

	/**
	 * GetByEmail retrieves a user by their email.
	 * @param ctx context.Context
	 * @param email string
	 * @return (*UserResponse, error)
	 */
	GetByEmail(ctx context.Context, email string) (result *account.UserResponse, err error)

	/**
	 * GetByUsername retrieves a user by their username.
	 * @param ctx context.Context
	 * @param username string
	 * @return (*UserResponse, error)
	 */
	GetByUsername(ctx context.Context, username string) (result *account.UserResponse, err error)

	/**
	 * Create creates a new user.
	 * @param ctx context.Context
	 * @param user *CreateUserRequest
	 * @return error
	 */
	Create(ctx context.Context, user *account.CreateUserRequest) (err error)

	/**
	 * Update updates a user.
	 * @param ctx context.Context
	 * @param id identity.ID
	 * @param user *UpdateUserRequest
	 * @return error
	 */
	Update(ctx context.Context, id identity.ID, user *account.UpdateUserRequest) (err error)

	/**
	 * Delete deletes a user.
	 * @param ctx context.Context
	 * @param id identity.ID
	 * @return error
	 */
	Delete(ctx context.Context, id identity.ID) (err error)
}
