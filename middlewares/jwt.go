package middlewares

import (
	"hashing-file/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func CreateToken(id string) (*string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    id,
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})
	token, err := claims.SignedString([]byte(config.LoadEnv().GetString("JWT_KEY")))
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func ClearToken() (*fiber.Cookie, error) {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	return &cookie, nil
}
