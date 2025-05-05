package interfaces

import (
	"context"

	"github.com/HasanNugroho/go-broilerplate-ddd/internal/domain/account"
	"github.com/HasanNugroho/go-broilerplate-ddd/internal/sharekernel/identity"
)

type IUserRepository interface {
	/**
	 * GetALl retrieves a user by filter.
	 * @param ctx context.Context
	 * @param limit int
	 * @param page int
	 * @param sort string
	 * @param search string
	 * @return (*UserResponse, error)
	 */
	GetAll(ctx context.Context, search string, limit int, page int, sort string) (result []*account.User, totalItems int64, err error)

	/**
	 * GetByID retrieves a user by their ID.
	 * @param ctx context.Context
	 * @param id identity.ID
	 * @return (*User, error)
	 */
	GetByID(ctx context.Context, id identity.ID) (result *account.User, err error)

	/**
	 * GetByEmail retrieves a user by their email.
	 * @param ctx context.Context
	 * @param email string
	 * @return (*User, error)
	 */
	GetByEmail(ctx context.Context, email string) (result *account.User, err error)

	/**
	 * GetByUsername retrieves a user by their username.
	 * @param ctx context.Context
	 * @param username string
	 * @return (*User, error)
	 */
	GetByUsername(ctx context.Context, username string) (result *account.User, err error)

	/**
	 * Create creates a new user.
	 * @param ctx context.Context
	 * @param user *User
	 * @return (*User, error)
	 */
	Create(ctx context.Context, user *account.User) (err error)

	/**
	 * Update updates a user.
	 * @param ctx context.Context
	 * @param user *User
	 * @return (*User, error)
	 */
	Update(ctx context.Context, user *account.User) (err error)

	/**
	 * Delete deletes a user.
	 * @param ctx context.Context
	 * @param id identity.ID
	 * @return error
	 */
	Delete(ctx context.Context, id identity.ID) (err error)
}
