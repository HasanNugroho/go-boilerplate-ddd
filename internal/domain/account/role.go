package account

import (
	"time"

	"github.com/HasanNugroho/go-broilerplate-ddd/internal/sharekernel/identity"
)

type Role struct {
	ID          identity.ID `json:"id" gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	Name        string      `json:"name" gorm:"column:name"`
	Permissions []string    `json:"permissions" gorm:"column:permissions"`
	CreatedAt   time.Time   `json:"created_at,omitempty" gorm:"column:created_at,omitempty"`
	UpdatedAt   time.Time   `json:"column:updated_at,omitempty" gorm:"column:updated_at,omitempty"`
}

func (Role) TableName() string {
	return "roles"
}
