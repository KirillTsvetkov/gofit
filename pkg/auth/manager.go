package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Manager struct {
	jwtKey string
}

func NewManager(jwtKey string) (*Manager, error) {
	if jwtKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{jwtKey: jwtKey}, nil
}

func (m *Manager) GenerateJWT(userID string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(ttl).Unix(),
	})

	tokenString, err := token.SignedString([]byte(m.jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (m *Manager) ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}
