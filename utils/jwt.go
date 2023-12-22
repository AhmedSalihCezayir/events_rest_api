package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64, isAdmin bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"userId":  userId,
		"isAdmin": isAdmin,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, bool, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, false, errors.New("Could not parse token!")
	}

	if !parsedToken.Valid {
		return 0, false, errors.New("Invalid token!")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, false, errors.New("Invalid token claims!")
	}

	// email := claims["email"].(string)
	id := int64(claims["userId"].(float64))
	isAdmin := claims["isAdmin"].(bool)
	return id, isAdmin, nil
}
