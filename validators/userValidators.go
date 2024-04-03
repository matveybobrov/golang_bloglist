package validators

import "bloglist/models"

type User = models.User

func ValidateUser(user User) (errorMessage, isValid) {
	if user.Username == "" {
		return "Username must be provided", false
	}
	if user.Name == "" {
		return "Name must be provided", false
	}
	if user.Password == "" {
		return "Password must be provided", false
	}
	return "", true
}
