package entity

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

type Role struct {
	ID   uuid.UUID `json:"id" validate:"required,uuid"`
	Name string    `json:"name" validate:"required,min=3"`
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

	role := &Role{
		ID:   id,
		Name: name,
	}

	if err = role.Validate(); err != nil {
		return nil, err
	}

	return role, nil
}
