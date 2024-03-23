package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/KirillTsvetkov/gofit/internal/domain"

	"github.com/dgrijalva/jwt-go"
)

type Manager struct {
	jwtKey string
}

type AuthClaims struct {
	jwt.StandardClaims
	User *domain.User `json:"user"`
}

func NewManager(jwtKey string) (*Manager, error) {
	if jwtKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{jwtKey: jwtKey}, nil
}

func (m *Manager) GenerateJWT(user *domain.User, ttl time.Duration) (string, error) {
	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(m.jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (m *Manager) ValidateJWT(tokenString string) (*domain.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.jwtKey), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, fmt.Errorf("invalid access token")
}
