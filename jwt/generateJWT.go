package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var signedKey = []byte("xl8gXatWkB6lr8rGzNqgsE29EzLsqCIb")

func generateJWTAcessToken(userID, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenString.SignedString(signedKey)
	if err != nil {
		return "", fmt.Errorf("error in generate JWT: %v", err)
	}

	return token, nil
}
