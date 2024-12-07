package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"my-gin-app/internal/models"
	"my-gin-app/internal/repositories"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepo}
}

func (s *AuthService) Signup(email, password string) error {
	// ハッシュ化されたパスワードを生成
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// ユーザーを作成
	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.userRepository.CreateUser(&user); err != nil {
		return errors.New("failed to create user")
	}

	return nil
}
