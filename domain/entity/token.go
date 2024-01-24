package entity

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidToken = errors.New("auth: invalid token")

type Token struct {
	UserUUId  string `json:"userUUId"`
	UserEmail string `json:"userEmail"`
	jwt.RegisteredClaims
}

func NewToken(userId string, userEmail string, duration time.Duration) *Token {
	return &Token{
		UserUUId:  userId,
		UserEmail: userEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

func (t *Token) SignToken(secret string) (tokenString string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, t)
	ss, err := token.SignedString([]byte(secret))
	return ss, err
}

func ParseToken(token string, secret string) (*Token, error) {

	parsedToken, err := jwt.ParseWithClaims(token, &Token{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	mapClaims, ok := parsedToken.Claims.(*Token)
	if !ok {
		return nil, ErrInvalidToken
	}

	return mapClaims, nil
}
