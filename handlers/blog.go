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
	blogs := []Blog{}
	rows, err := db.DB.Query("SELECT * FROM blogs")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	for rows.Next() {
		blog := Blog{}
		err := rows.Scan(&blog.Id, &blog.Title, &blog.Author, &blog.Url, &blog.Likes)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		blogs = append(blogs, blog)
	}

	json.NewEncoder(res).Encode(blogs)
}

func GetOneBlog(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil || id < 0 {
		http.Error(res, "Malformatted id", 400)
		return
	}

	blog := Blog{}
	row := db.DB.QueryRow("SELECT * FROM blogs WHERE id=$1", id)
	err = row.Scan(&blog.Id, &blog.Author, &blog.Url, &blog.Title, &blog.Likes)
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

	savedBlog := Blog{}
	row := db.DB.QueryRow("INSERT INTO blogs (title, author, url) VALUES ($1, $2, $3) RETURNING *", blog.Title, blog.Author, blog.Url)
	err := row.Scan(&savedBlog.Id, &savedBlog.Title, &savedBlog.Author, &savedBlog.Url, &savedBlog.Likes)
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

	_, err = db.DB.Exec("DELETE FROM blogs WHERE id=$1", id)
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

	updatedBlog := Blog{}
	row := db.DB.QueryRow("UPDATE blogs SET title=$1, author=$2, url=$3 RETURNING *", blog.Title, blog.Author, blog.Url)
	err = row.Scan(&updatedBlog.Id, &updatedBlog.Title, &updatedBlog.Author, &updatedBlog.Url, &updatedBlog.Likes)
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
