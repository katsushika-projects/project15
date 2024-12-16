package middlewares

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/moto340/project15/backend/internal/repositories"
)

type AuthMiddleware struct {
	userRepository *repositories.UserRepository
}

func NewAuthMiddleware(userRepo *repositories.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{userRepository: userRepo}
}

var accessSecret = []byte(os.Getenv("ACCESS_SECRET_KEY"))

func parseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return accessSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (m *AuthMiddleware) AuthAccessToken(access_token string) error {
	if access_token == "" {
		return errors.New("access token does not exist")
	}

	tokenString := strings.TrimPrefix(access_token, "Bearer ")
	if tokenString == access_token {
		return errors.New("invalid token format")
	}

	// トークンがブラックリストにあるか確認
	if err := m.userRepository.AuthBlackList(tokenString); err != nil {
		return err
	}

	// トークンを解析・検証
	payload, err := parseJWT(tokenString)
	if err != nil {
		return err
	}

	// expの検証
	exp, ok := payload["exp"].(float64)
	if !ok {
		return errors.New("exp claim missing or not a float64")
	}
	if int64(exp) < time.Now().Unix() {
		return errors.New("token has expired")
	}

	return nil
}

/*
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

func (m *AuthMiddleware) AuthAccessToken(access_token string) error {
	if access_token == "" {
		return errors.New("access_token doesn't exist")
	}

	token_string := strings.TrimPrefix(access_token, "Bearer ")
	if token_string == access_token {
		return errors.New("invalid token format")
	}

	if err := m.userRepository.AuthBlackList(token_string); err != nil {
		return err
	}

	payload, err1 := parseJWT(token_string)
	if err1 != nil {
		return err1
	}

	exp, ok := payload["exp"].(float64)
	if !ok {
		return errors.New("exp claim missing or not a float64")
	}

	if int64(exp) < time.Now().Unix() {
		return errors.New("token has expire")
	}

	pre_access_token, err2 := jwt.NewWithClaims(jwt.SigningMethodES256, payload).SignedString(accessSecret)
	if err2 != nil {
		return err2
	}

	if access_token != pre_access_token {
		return errors.New("invalid token")
	}

	return nil
}
*/
