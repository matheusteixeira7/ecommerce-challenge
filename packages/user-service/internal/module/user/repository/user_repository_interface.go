package repository

import (
	"github.com/google/uuid"
	"github.com/matheusteixeira7/ecommerce-challenge/packages/user-service/internal/module/user/entity"
)

type UserRepositoryInterface interface {
	Save(user *entity.User) (uuid.UUID, error)
	FindByID(id uuid.UUID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
