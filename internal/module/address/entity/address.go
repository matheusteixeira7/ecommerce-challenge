package entity

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

type Address struct {
	Id      uuid.UUID `json:"id"`
	UserId  uuid.UUID `json:"user_id" validate:"required"`
	Street  string    `json:"street" validate:"required,min=3"`
	City    string    `json:"city" validate:"required,min=2"`
	State   string    `json:"state" validate:"required,len=2"`
	ZipCode string    `json:"zip_code" validate:"required,len=8,numeric"`
}

func (a *Address) Validate() error {
	err := validate.Struct(a)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return errors
	}
	return nil
}

func NewAddress(userId uuid.UUID, street, city, state, zipCode string) (*Address, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	address := &Address{
		Id:      id,
		UserId:  userId,
		Street:  street,
		City:    city,
		State:   state,
		ZipCode: zipCode,
	}

	if err := address.Validate(); err != nil {
		fmt.Println("Validation errors:", err)
		return nil, err
	}

	return address, nil
}
