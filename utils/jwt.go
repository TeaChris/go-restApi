package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "secret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte (secret))
}


func VerifyToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("invalid token")
	}

	if !parsedToken.Valid {
		return nil, errors.New("token is invalid")
	}

	return parsedToken, nil
}
