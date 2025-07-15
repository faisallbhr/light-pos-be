package jwtx

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type JWTManager struct {
	accessSecret  string
	refreshSecret string
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJWTManager(accessSecret, refreshSecret string, accessTTL, refreshTTL time.Duration) *JWTManager {
	return &JWTManager{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessTTL:     accessTTL,
		refreshTTL:    refreshTTL,
	}
}

func (j *JWTManager) GenerateToken(userID uint, tokenType TokenType) (string, error) {
	var secret string
	var ttl time.Duration

	switch tokenType {
	case AccessToken:
		secret = j.accessSecret
		ttl = j.accessTTL
	case RefreshToken:
		secret = j.refreshSecret
		ttl = j.refreshTTL
	default:
		return "", errors.New("invalid token type")
	}

	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   string(tokenType),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (j *JWTManager) ValidateToken(tokenString string, tokenType TokenType) (*CustomClaims, error) {
	var secret string

	switch tokenType {
	case AccessToken:
		secret = j.accessSecret
	case RefreshToken:
		secret = j.refreshSecret
	default:
		return nil, errors.New("invalid token type")
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
