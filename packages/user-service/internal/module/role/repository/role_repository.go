package repository

import (
	"database/sql"
	"github.com/matheusteixeira7/ecommerce-challenge/packages/user-service/internal/module/role/entity"
)

type RoleRepository struct {
	db *sql.DB
}

func (r *RoleRepository) FindMany() ([]*entity.Role, error) {
	query := `SELECT id, name, created_at, updated_at FROM roles`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			err = closeErr
		}
	}()

	var roles []*entity.Role

	for rows.Next() {
		var role entity.Role
		if err := rows.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}
