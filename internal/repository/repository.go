package repository

import (
	"todoapp/internal/entity"

	"github.com/jmoiron/sqlx"
)

type IAuthorization interface {
	CreateUser(user entity.User) (int, error)
}

type Repository struct {
	IAuthorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IAuthorization: NewAuthPostgres(db),
	}
}
