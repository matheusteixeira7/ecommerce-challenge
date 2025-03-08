package entity

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

var validate = validator.New()

type User struct {
	ID        uuid.UUID  `json:"id" validate:"required,uuid"`
	Name      string     `json:"name" validate:"required,min=3,max=100"`
	Email     string     `json:"email" validate:"required,email"`
	Password  string     `json:"password" validate:"required,gte=8"`
	RoleID    uuid.UUID  `json:"role_id" validate:"required,uuid"`
	CreatedAt time.Time  `json:"created_at" validate:"required"`
	UpdatedAt time.Time  `json:"updated_at" validate:"required"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (u *User) Validate() error {
	err := validate.Struct(u)
	if err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}

func NewUser(name, email, password string, roleID uuid.UUID) (*User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	user := &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		RoleID:    roleID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err = user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}
