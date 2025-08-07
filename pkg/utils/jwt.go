package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/daytrip-idn-api/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateSessionToken(user *models.User) (string, error) {
	hoursCount, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOURS"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["user_id"] = user.UserId
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyToken(tokenString string, isEmployee bool) (*jwt.Token, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else {
				return nil, errors.New("token is invalid")
			}
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	return token, nil
}

func ExtractClaims(token *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok && token.Valid
}
