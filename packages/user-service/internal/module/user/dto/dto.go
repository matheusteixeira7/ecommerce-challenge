package dto

import "github.com/google/uuid"

type CreateUserInputDTO struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     uuid.UUID `json:"role"`
}

type CreateUserOutputDTO struct {
	ID uuid.UUID `json:"id"`
}
