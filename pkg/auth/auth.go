package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = os.Getenv("SECRET_KEY")

type CustomClaims struct {
	Email string
	jwt.RegisteredClaims
}

func GenerateToken(email string) (string, error) {
	claims := &CustomClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return ss, nil
}

func ValidateToken(ss string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(ss, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok && !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
