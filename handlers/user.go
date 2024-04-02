package handlers

import (
	"bloglist/db"
	"bloglist/helpers"
	"bloglist/models"
	"bloglist/validators"
	"encoding/json"
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
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
