package handlers

import (
	"bloglist/db"
	"bloglist/helpers"
	"bloglist/models"
	"bloglist/validators"
	"encoding/json"
	"fmt"
	"net/http"
)

type User = models.User
type UserWithToken = models.UserWithToken

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)

	msg, ok := validators.ValidateUser(user)
	if !ok {
		http.Error(w, msg, 400)
		return
	}

	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	user.Password = hash

	savedUser, err := db.InsertUser(user)
	if err != nil {
		// TODO: handle unique error properly
		http.Error(w, err.Error(), 500)
		return
	}

	token, err := helpers.SignToken(savedUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := UserWithToken{
		User:  savedUser,
		Token: token,
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(response)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)

	foundUser, err := db.GetUserByUsername(user.Username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Println(user.Password, foundUser.Password)
	ok := helpers.CheckPasswordHash(user.Password, foundUser.Password)
	if !ok {
		http.Error(w, "Incorrect password or username", http.StatusUnauthorized)
		return
	}

	token, err := helpers.SignToken(foundUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Hide password from response
	foundUser.Password = ""
	response := UserWithToken{
		User:  foundUser,
		Token: token,
	}

	json.NewEncoder(w).Encode(response)
}
