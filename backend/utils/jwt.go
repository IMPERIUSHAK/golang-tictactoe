package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var Jwtsecret []byte

type UserClaims struct {
	UserId   int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(userid int, username string) (string, error) {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	Jwtsecret = []byte(os.Getenv("SECRET"))

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := UserClaims{
		UserId:   userid,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(Jwtsecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (*UserClaims, error) {
	claims := &UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return Jwtsecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
