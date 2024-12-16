package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (s *AuthService) Login(username, password string) error {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return errors.New("username doesn't exist")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("password doesn't match")
	}

	return nil
}

var accessSecret = []byte(os.Getenv("ACCESS_SECRET_KEY"))
var refreshSecret = []byte(os.Getenv("REFRESH_SECRET_KEY"))

func (s *AuthService) GenerateTokens(username string) (accessToken, refreshToken string, err error) {
	user, _ := s.userRepository.FindByUsername(username)
	accessClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(accessSecret)
	if err != nil {
		return "", "", err
	}

	// リフレッシュトークン
	refreshClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 有効期限7日
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(refreshSecret)
	if err != nil {
		return "", "", err
	}

	err = s.userRepository.CreateAccessToken(user, refreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func parseJWT(token string) (jwt.MapClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("Invalid token format")
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	var payload jwt.MapClaims
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, err
	}

	return payload, nil
}

func (s *AuthService) RefreshTokenDisable(token string) error {
	token_string := strings.TrimPrefix(token, "Bearer ")

	payload, err1 := parseJWT(token_string)
	if err1 != nil {
		return err1
	}

	user_id := payload["user_id"].(string)
	if err := s.userRepository.UpdateRefreshToken(user_id); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) AccessTokenDisable(access_token string) error {
	// アクセストークンをブラックリストに入れる
	if err := s.userRepository.PostBlackList(access_token); err != nil {
		return err
	}

	return nil
}
