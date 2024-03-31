package main

import (
	"bloglist/db"
	"bloglist/handlers"

	"log"
	"net/http"
)

func init() {
	log.Println("Connecting to database...")
	err := db.Init()
	if err != nil {
		log.Fatal("Failed connecting to database\n", err.Error())
	}
	log.Println("Connected to database")
}

// TODO: make centrilized error handler
// TODO: make authorization middleware
// TODO: add users
func main() {
	http.HandleFunc("GET /api/blogs", handlers.GetAllBlogs)
	http.HandleFunc("GET /api/blogs/{id}", handlers.GetOneBlog)
	http.HandleFunc("POST /api/blogs", handlers.CreateOneBlog)
	http.HandleFunc("DELETE /api/blogs/{id}", handlers.DeleteOneBlog)
	http.HandleFunc("PUT /api/blogs/{id}", handlers.UpdateOneBlog)

	log.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
