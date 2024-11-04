package repository

import (
	"fmt"
	"time"
	"todoapp/internal/entity"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int

	query := fmt.Sprintf(`
		INSERT INTO %s (username, email, password, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id`, usersTable)

	now := time.Now()

	row := r.db.QueryRow(query, user.Username, user.Email, user.Password, now, now)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
