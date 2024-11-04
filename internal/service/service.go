package service

import (
	"todoapp/internal/entity"
	"todoapp/internal/repository"
)

type IAuthorization interface {
	CreateUser(user entity.User) (int, error)
}

type Service struct {
	IAuthorization
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		IAuthorization: NewAuthService(rep.IAuthorization),
	}
}
