package services

import (
	"grafolito/backend/internal/identity/domain/entities"
	"grafolito/backend/internal/shared/application/options"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user *entities.User) (string, error) {
	jwtOptions := options.JwtOptions()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"iss": jwtOptions.Issuer,
		"aud": jwtOptions.Audience,
	})
	jwt, err := token.SignedString([]byte(jwtOptions.Secret))
	if err != nil {
		return "", err
	}

	return jwt, nil
}
