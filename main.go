package main

import (
	"bloglist/db"
	"bloglist/handlers"

	"fmt"
	"net/http"
	"os"
)

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
	http.HandleFunc("GET /api/blogs", handlers.GetAllBlogs)
	http.HandleFunc("GET /api/blogs/{id}", handlers.GetOneBlog)
	http.HandleFunc("POST /api/blogs", handlers.CreateOneBlog)
	http.HandleFunc("DELETE /api/blogs/{id}", handlers.UpdateOneBlog)

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
