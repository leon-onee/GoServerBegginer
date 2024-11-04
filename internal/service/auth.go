package service

import (
	"todoapp/internal/entity"
	"todoapp/internal/repository"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	rep repository.IAuthorization
}

func NewAuthService(rep repository.IAuthorization) *AuthService {
	return &AuthService{
		rep: rep,
	}
}

func (s *AuthService) CreateUser(user entity.User) (int, error) {
	hashedPassword, err := generateHashPassword(user.Password)
	if err != nil {
		logrus.Fatalf("Ошибка шифрования пароля: %s", err)
	}
	user.Password = hashedPassword
	return s.rep.CreateUser(user)
}

func generateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
