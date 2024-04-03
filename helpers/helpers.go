package helpers

import (
	"bloglist/models"
	"os"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User = models.User

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignToken(user User) (string, error) {
	// create token with some data (map)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"name":     user.Name,
	})
	secret := os.Getenv("TOKEN_SECRET")
	// sign token with secret
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func ParseToken(token string) (User, error) {
	user := User{}
	secret := os.Getenv("TOKEN_SECRET")

	parsedToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return user, err
	}

	claims, _ := parsedToken.Claims.(jwt.MapClaims)
	user.Id = int(claims["id"].(float64))
	user.Username = claims["username"].(string)
	user.Name = claims["name"].(string)

	return user, nil
}
