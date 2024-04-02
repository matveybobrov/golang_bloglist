package helpers

import (
	"bloglist/models"
	"os"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User = models.User

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignToken(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
	})
	secret := os.Getenv("TOKEN_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
