package account

import (
	"time"

	"github.com/HasanNugroho/go-broilerplate-ddd/internal/sharekernel/identity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        identity.ID `json:"id" gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	Name      string      `json:"name" gorm:"column:name"`
	Fullname  string      `json:"fullname" gorm:"column:fullname"`
	Username  string      `json:"username" gorm:"column:username;unique"`
	Email     string      `json:"email" gorm:"column:email"`
	Password  string      `json:"password" gorm:"column:password"`
	Role      identity.ID `json:"role_id" gorm:"column:role_id;type:uuid;"`
	RoleData  Role        `json:"role_data" gorm:"-"`
	IsActive  bool        `json:"is_active" gorm:"column:is_active;default:true"`
	CreatedAt time.Time   `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time   `json:"updated_at" gorm:"column:updated_at"`
}

type (
	UserResponse struct {
		ID        identity.ID `json:"id"`
		Email     string      `json:"email"`
		Name      string      `json:"name"`
		Role      Role        `json:"role"`
		CreatedAt time.Time   `json:"created_at,omitempty"`
		UpdatedAt time.Time   `json:"updated_at,omitempty"`
	}

	CreateUserRequest struct {
		Name     string `json:"name" validate:"required"`
		Fullname string `json:"fullname" validate:"required"`
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}

	UpdateUserRequest struct {
		Name     string      `json:"name" validate:"name"`
		Fullname string      `json:"fullname" validate:"fullname"`
		Username string      `json:"username" validate:"username"`
		Email    string      `json:"email" validate:"email"`
		Password string      `json:"password" validate:"password"`
		Role     identity.ID `json:"role_id" gorm:"column:role_id;type:uuid;"`
	}
)

func (User) TableName() string {
	return "users"
}

func (u *User) EncryptPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) VerifyPassword(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))
	return err == nil
}

func (u *User) ToUserResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		Role:      u.RoleData,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
