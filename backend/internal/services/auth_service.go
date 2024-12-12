package services

import (
	"errors"

	"github.com/moto340/project15/backend/internal/models"
	"github.com/moto340/project15/backend/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepo}
}

func (s *AuthService) Signup(username, password string) error {
	// ユーザー名の重複確認
	if _, err := s.userRepository.FindByUsername(username); err == nil {
		return errors.New("username already exists")
	}

	// ハッシュ化されたパスワードを生成
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// ユーザーを作成
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	if err := s.userRepository.CreateUser(&user); err != nil {
		return errors.New("failed to create user")
	}

	return nil
}
