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

func GetAllBlogs(res http.ResponseWriter, req *http.Request) {
	blogs, err := db.GetAllBlogs()
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	json.NewEncoder(res).Encode(blogs)
}

func GetOneBlog(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil || id < 0 {
		http.Error(res, "Malformatted id", 400)
		return
	}

	blog, err := db.GetOneBlog(id)
	if err == sql.ErrNoRows {
		message := fmt.Sprintf("Blog with id %v was not found", id)
		http.Error(res, message, 404)
		return
	}
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	json.NewEncoder(res).Encode(blog)
}

func CreateOneBlog(res http.ResponseWriter, req *http.Request) {
	blog := Blog{}
	json.NewDecoder(req.Body).Decode(&blog)

	msg, ok := validators.ValidateBlog(blog)
	if !ok {
		http.Error(res, msg, 400)
		return
	}

	savedBlog, err := db.CreateOneBlog(blog)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	res.WriteHeader(201)
	json.NewEncoder(res).Encode(savedBlog)
}

func DeleteOneBlog(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil || id < 0 {
		http.Error(res, "Malformatted id", 400)
		return
	}

	err = db.DeleteOneBlog(id)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	res.WriteHeader(204)
}

func UpdateOneBlog(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil || id < 0 {
		http.Error(res, "Malformatted id", 400)
		return
	}

	blog := Blog{}
	json.NewDecoder(req.Body).Decode(&blog)

	msg, ok := validators.ValidateBlog(blog)
	if !ok {
		http.Error(res, msg, 400)
		return
	}

	updatedBlog, err := db.UpdateOneBlog(blog, id)
	if err == sql.ErrNoRows {
		message := fmt.Sprintf("Blog with id %v was not found", id)
		http.Error(res, message, 404)
		return
	}
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	json.NewEncoder(res).Encode(updatedBlog)
}
