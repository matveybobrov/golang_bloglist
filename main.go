package main

import (
	"bloglist/db"
	"bloglist/handlers"
	"bloglist/middlewares"

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
	Logger := middlewares.Logger

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/blogs", handlers.GetAllBlogs)
	mux.HandleFunc("GET /api/blogs/{id}", handlers.GetOneBlog)
	mux.HandleFunc("POST /api/blogs", handlers.CreateOneBlog)
	mux.HandleFunc("DELETE /api/blogs/{id}", handlers.DeleteOneBlog)
	mux.HandleFunc("PUT /api/blogs/{id}", handlers.UpdateOneBlog)

	server := Logger(mux)

	log.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", server)
}
