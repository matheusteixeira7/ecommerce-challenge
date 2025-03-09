package repository

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/matheusteixeira7/ecommerce-challenge/packages/user-service/internal/module/user/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(user *entity.User) (uuid.UUID, error) {
	query := `INSERT INTO users (id, name, email, password, role_id, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := r.db.QueryRow(query, user.ID, user.Name, user.Email, user.Password, user.RoleID, user.CreatedAt, user.UpdatedAt, user.DeletedAt).Scan(&user.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	query := `SELECT u.id, u.name, u.email, u.password, u.role_id, u.created_at, u.updated_at, u.deleted_at 
	          FROM users u 
	          WHERE u.email = $1`

	var user entity.User
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
