package repository

import (
	"github.com/google/uuid"
	"github.com/matheusteixeira7/ecommerce-challenge/packages/user-service/internal/module/role/entity"
)

type RoleRepositoryInterface interface {
	FindMany() ([]*entity.Role, error)
	FindUniqueByID(id uuid.UUID) ([]*entity.Role, error)
	FindUniqueByName(name string) ([]*entity.Role, error)
}
