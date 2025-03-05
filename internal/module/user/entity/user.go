package entity

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	RoleId   string    `json:"role_id"`
}

func NewUser(name, email, password, roleId string) (*User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	user := &User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
		RoleId:   roleId,
	}

	return user, nil
}
