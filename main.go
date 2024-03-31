package main

import (
	"bloglist/db"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Blog struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Url    string `json:"url"`
	Likes  int    `json:"likes"`
}

func init() {
	fmt.Println("Connecting to database")
	err := db.Init()
	if err != nil {
		fmt.Println("Failed connecting to database")
		os.Exit(1)
	}
	fmt.Println("Connected to database")
}

func main() {
	http.HandleFunc("GET /api/blogs", func(res http.ResponseWriter, req *http.Request) {
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
	})

	http.HandleFunc("GET /api/blogs/{id}", func(res http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(req.PathValue("id"))
		if err != nil || id < 0 {
			http.Error(res, "Malformatted id", 400)
			return
		}

		blog := Blog{}
		row := db.DB.QueryRow("SELECT * FROM blogs WHERE id=$1", id)
		err = row.Scan(&blog.Id, &blog.Author, &blog.Url, &blog.Title, &blog.Likes)
		if err == sql.ErrNoRows {
			message := fmt.Sprintf("Person with id %v was not found", id)
			http.Error(res, message, 404)
			return
		} else if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		json.NewEncoder(res).Encode(blog)
	})

	http.HandleFunc("POST /api/blogs", func(res http.ResponseWriter, req *http.Request) {
		blog := Blog{}
		json.NewDecoder(req.Body).Decode(&blog)

		if blog.Title == "" {
			http.Error(res, "Title must be provided", 400)
			return
		}
		if blog.Author == "" {
			http.Error(res, "Author must be provided", 400)
			return
		}
		if blog.Url == "" {
			http.Error(res, "Url must be provided", 400)
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
	})

	http.HandleFunc("DELETE /api/blogs/{id}", func(res http.ResponseWriter, req *http.Request) {
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
	})

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
