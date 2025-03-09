package entity

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

var validate = validator.New()

type Role struct {
	ID        uuid.UUID `json:"id" validate:"required,uuid"`
	Name      string    `json:"name" validate:"required,min=3"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}

func (r *Role) Validate() error {
	err := validate.Struct(r)
	if err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}

func NewRole(name string) (*Role, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	role := &Role{
		ID:        id,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err = role.Validate(); err != nil {
		return nil, err
	}

	return role, nil
}
