package middlewares

import (
	"hashing-file/config"

	"github.com/golang-jwt/jwt"
)

func Auth(cookie string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.LoadEnv().GetString("JWT_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractAuth(token *jwt.Token) *jwt.StandardClaims {
	claims := token.Claims.(*jwt.StandardClaims)

	return claims
}
