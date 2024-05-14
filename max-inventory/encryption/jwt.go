package encryption

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jorgeemherrera/Golang/internal/models"
)

func SignedLoginToken(user *models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
	})

	jwtString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return jwtString, nil
}
