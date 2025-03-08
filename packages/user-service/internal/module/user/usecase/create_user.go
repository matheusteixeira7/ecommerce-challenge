package usecase

import (
	"errors"
	"github.com/matheusteixeira7/ecommerce-challenge/packages/user-service/internal/module/user/dto"
	"github.com/matheusteixeira7/ecommerce-challenge/packages/user-service/internal/module/user/entity"
	"github.com/matheusteixeira7/ecommerce-challenge/packages/user-service/internal/module/user/repository"
)

type UserRepository struct {
	repository repository.UserRepositoryInterface
}

func NewCreateUserUseCase(repository repository.UserRepositoryInterface) *UserRepository {
	return &UserRepository{
		repository: repository,
	}
}

func (r *UserRepository) Execute(input dto.CreateUserInputDTO) (*dto.CreateUserOutputDTO, error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password, input.Role)
	if err != nil {
		return nil, err
	}

	existingUser, err := r.repository.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	userID, err := r.repository.Save(user)
	if err != nil {
		return nil, err
	}
	return &dto.CreateUserOutputDTO{
		ID: userID,
	}, nil
}
