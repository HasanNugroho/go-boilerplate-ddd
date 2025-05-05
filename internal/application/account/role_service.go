package account

import (
	"context"
	"time"

	"github.com/HasanNugroho/go-broilerplate-ddd/internal/domain/account"
	"github.com/HasanNugroho/go-broilerplate-ddd/internal/domain/account/interfaces"
	"github.com/HasanNugroho/go-broilerplate-ddd/internal/sharekernel/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type RoleService struct {
	repo interfaces.IRoleRepository
}

func NewRoleService(repo interfaces.IRoleRepository) *RoleService {
	return &RoleService{
		repo: repo,
	}
}

func (r *RoleService) Create(ctx context.Context, role *account.Role) (err error) {
	payload := account.Role{
		Name:        role.Name,
		Permissions: role.Permissions,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := r.repo.Create(ctx, &payload); err != nil {
		return err
	}

	return nil
}

func (r *RoleService) FindById(ctx context.Context, id string) (result *account.Role, err error) {
	role, err := r.repo.FindById(ctx, id)
	if err != nil {
		return &account.Role{}, err
	}
	return role, err
}

func (r *RoleService) FindManyByID(ctx context.Context, ids []bson.ObjectID) (result *[]account.Role, err error) {
	roles, err := r.repo.FindManyByID(ctx, ids)
	if err != nil {
		return &[]account.Role{}, err
	}

	return roles, nil
}

func (r *RoleService) FindAll(ctx context.Context, filter *model.PaginationFilter) (result *[]account.Role, totalItems int64, err error) {
	roles, totalItems, err := r.repo.FindAll(ctx, filter)
	if err != nil {
		return &[]account.Role{}, 0, err
	}

	return roles, int64(totalItems), nil
}

func (r *RoleService) Update(ctx context.Context, id string, role *account.Role) (err error) {
	currentRole, err := r.repo.FindById(ctx, id)
	if err != nil {
		return err
	}

	if role.Name != "" {
		currentRole.Name = role.Name
	}

	return r.repo.Update(ctx, id, currentRole)
}

func (r *RoleService) Delete(ctx context.Context, id string) (err error) {
	err = r.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoleService) AssignUser(ctx context.Context, userId string, roleId string) (err error) {
	err = r.repo.AssignUser(ctx, userId, roleId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoleService) UnassignUser(ctx context.Context, userId string, roleId string) (err error) {
	err = r.repo.UnassignUser(ctx, userId, roleId)
	if err != nil {
		return err
	}
	return nil
}
