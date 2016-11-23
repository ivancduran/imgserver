package libs

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWT struct {
	UserName string
}

func (j JWT) Secret() []byte {
	return []byte("ZA1XSWSekret128cdevfraASDFlkjhHg")
}

func (j JWT) Create(email string, username string, role int8) *jwt.Token {

	// time := time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix()

	now := time.Now()
	secs := now.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"username": username,
		"role":     role,
		"nbf":      secs,
	})

	return token
}
