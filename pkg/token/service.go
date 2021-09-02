package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Subject string `json:"sub,omitempty"`
	Role    string `json:"role,omitempty"`
}

type ClaimsResult struct {
	jwt.StandardClaims
	Role string `json:"role,omitempty"`
}

type Service interface {
	GenerateToken(claims *Claims, expiration int) (string, error)
	ValidateToken(tokenString string) (*ClaimsResult, error)
}

type service struct {
	Secret []byte
}

func NewJWTService(secret []byte) Service {
	return &service{
		Secret: secret,
	}
}

func (service *service) GenerateToken(claims *Claims, expiration int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  claims.Subject,
		"role": claims.Role,
		"exp":  time.Now().Add(time.Hour * time.Duration(expiration)).Unix(),
	})
	tokenString, err := token.SignedString(service.Secret)
	return tokenString, err
}

func (service *service) ValidateToken(tokenString string) (*ClaimsResult, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ClaimsResult{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return service.Secret, nil
	})
	if claims, ok := token.Claims.(*ClaimsResult); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
