package presistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/HasanNugroho/go-broilerplate-ddd/internal/domain/account"
	"github.com/HasanNugroho/go-broilerplate-ddd/internal/sharekernel/errs"
	"github.com/HasanNugroho/go-broilerplate-ddd/internal/sharekernel/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) Create(ctx context.Context, role *account.Role) (err error) {
	panic("not implemented") // TODO: Implement
}

func (r *RoleRepository) FindById(ctx context.Context, id string) (result *account.Role, err error) {
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("role with ID '%s' not found: %w", id, errs.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to query role: %w", err)
	}
	return result, nil
}

func (r *RoleRepository) FindManyByID(ctx context.Context, ids []bson.ObjectID) (result *[]account.Role, err error) {
	if err := r.db.WithContext(ctx).Where("id IN (?)", ids).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("role not found: %w", errs.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to query role: %w", err)
	}
	return result, nil
}

func (r *RoleRepository) FindAll(ctx context.Context, filter *model.PaginationFilter) (result *[]account.Role, totalItems int64, err error) {
	offset := (filter.Page - 1) * filter.Limit

	sortField := "updated_at"
	if filter.Sort != "asc" && filter.Sort != "desc" {
		filter.Sort = "asc"
	}
	orderBy := fmt.Sprintf("%s %s", sortField, filter.Sort)

	query := r.db.WithContext(ctx).Model(&result).Order(orderBy).Limit(filter.Limit).Offset(offset)

	if filter.Search != "" {
		searchPattern := fmt.Sprintf("%%%s%%", filter.Search)
		query = query.Where("name ILIKE ? ", searchPattern, searchPattern)
	}

	if err = query.Find(&result).Error; err != nil {
		return nil, 0, fmt.Errorf("role not found: %w", errs.ErrNotFound)
	}

	if err := query.Model(result).Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (r *RoleRepository) Update(ctx context.Context, id string, role *account.Role) (err error) {
	result := r.db.WithContext(ctx).Save(role)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("role with ID '%s' not found: %w", role.ID, errs.ErrNotFound)
		}
		return fmt.Errorf("failed to update role: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("role with ID '%s' not found: %w", role.ID, errs.ErrNotFound)
	}
	return nil
}

func (r *RoleRepository) Delete(ctx context.Context, id string) (err error) {
	result := r.db.WithContext(ctx).Where("id", id).Delete(&account.Role{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("role with ID '%s' not found: %w", id, errs.ErrNotFound)
		}
		return fmt.Errorf("failed to delete role: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("role with ID '%s' not found: %w", id, errs.ErrNotFound)
	}
	return nil
}

func (r *RoleRepository) AssignUser(ctx context.Context, userId string, roleId string) (err error) {
	panic("not implemented") // TODO: Implement
}

func (r *RoleRepository) UnassignUser(ctx context.Context, userId string, roleId string) (err error) {
	panic("not implemented") // TODO: Implement
}
