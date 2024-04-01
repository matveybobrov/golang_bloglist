package handlers

import (
	"bloglist/db"
	"bloglist/models"
	"bloglist/validators"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Blog = models.Blog

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := db.GetAllBlogs()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(blogs)
}

func GetOneBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.Error(w, "Malformatted id", 400)
		return
	}

	blog, err := db.GetOneBlog(id)
	if err == sql.ErrNoRows {
		message := fmt.Sprintf("Blog with id %v was not found", id)
		http.Error(w, message, 404)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(blog)
}

func CreateOneBlog(w http.ResponseWriter, r *http.Request) {
	blog := Blog{}
	json.NewDecoder(r.Body).Decode(&blog)

	msg, ok := validators.ValidateBlog(blog)
	if !ok {
		http.Error(w, msg, 400)
		return
	}

	savedBlog, err := db.CreateOneBlog(blog)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(savedBlog)
}

func DeleteOneBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.Error(w, "Malformatted id", 400)
		return
	}

	err = db.DeleteOneBlog(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(204)
}

func UpdateOneBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.Error(w, "Malformatted id", 400)
		return
	}

	blog := Blog{}
	json.NewDecoder(r.Body).Decode(&blog)

	msg, ok := validators.ValidateBlog(blog)
	if !ok {
		http.Error(w, msg, 400)
		return
	}

	updatedBlog, err := db.UpdateOneBlog(blog, id)
	if err == sql.ErrNoRows {
		message := fmt.Sprintf("Blog with id %v was not found", id)
		http.Error(w, message, 404)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(updatedBlog)
}
