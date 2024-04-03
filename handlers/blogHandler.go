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

// TODO: also return its user
func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := db.GetAllBlogs()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(blogs)
}

// TODO: also return its user
func GetOneBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.Error(w, "Malformatted id", 400)
		return
	}

	blog, err := db.GetBlogById(id)
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

	user := r.Context().Value("User").(User)
	blog.User_id = user.Id

	savedBlog, err := db.InsertBlog(blog)
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

	foundBlog, err := db.GetBlogById(id)
	if err == sql.ErrNoRows {
		message := fmt.Sprintf("Blog with id %v was not found", id)
		http.Error(w, message, 404)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user := r.Context().Value("User").(User)
	if user.Id != foundBlog.User_id {
		http.Error(w, "Not the creator of the blog", http.StatusForbidden)
		return
	}

	err = db.DeleteBlogById(id)
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

	foundBlog, err := db.GetBlogById(id)
	if err == sql.ErrNoRows {
		message := fmt.Sprintf("Blog with id %v was not found", id)
		http.Error(w, message, 404)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user := r.Context().Value("User").(User)
	if user.Id != foundBlog.User_id {
		http.Error(w, "Not the creator of the blog", http.StatusForbidden)
		return
	}

	updatedBlog, err := db.UpdateBlogById(blog, id)
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
