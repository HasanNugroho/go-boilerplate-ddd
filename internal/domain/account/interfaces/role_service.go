package interfaces

import (
	"context"

	"github.com/HasanNugroho/go-broilerplate-ddd/internal/domain/account"
	"github.com/HasanNugroho/go-broilerplate-ddd/internal/sharekernel/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type IRoleService interface {
	Create(ctx context.Context, role *account.Role) (err error)
	FindById(ctx context.Context, id string) (result *account.Role, err error)
	FindManyByID(ctx context.Context, ids []bson.ObjectID) (result *[]account.Role, err error)
	FindAll(ctx context.Context, filter *model.PaginationFilter) (result *[]account.Role, totalItems int64, err error)
	Update(ctx context.Context, id string, role *account.Role) (err error)
	Delete(ctx context.Context, id string) (err error)
	AssignUser(ctx context.Context, userId string, roleId string) (err error)
	UnassignUser(ctx context.Context, userId string, roleId string) (err error)
}
