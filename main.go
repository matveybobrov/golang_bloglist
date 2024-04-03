package main

import (
	"bloglist/db"
	"bloglist/handlers"
	"bloglist/middlewares"

	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	// TODO: maybe handle an error
	godotenv.Load()

	// log also prints date unlike fmt
	log.Println("Connecting to database...")
	err := db.Init()
	if err != nil {
		log.Fatal("Failed connecting to database\n", err.Error())
	}
	log.Println("Connected to database")
}

// TODO: consider sqlx
// TODO: change db columns order (esp. for users)
// TODO: work with queries
// TODO: make centrilized error handler
// TODO: generate swagger docs
func main() {
	Logger := middlewares.Logger
	UserExtractor := middlewares.UserExtractor

	// I had to create it in order to use global Logger
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/blogs", handlers.GetAllBlogsWithUsers)
	mux.HandleFunc("GET /api/blogs/{id}", handlers.GetOneBlogWithUser)
	mux.HandleFunc("POST /api/blogs", UserExtractor(handlers.CreateOneBlog))
	mux.HandleFunc("DELETE /api/blogs/{id}", UserExtractor(handlers.DeleteOneBlog))
	mux.HandleFunc("PUT /api/blogs/{id}", UserExtractor(handlers.UpdateOneBlog))

	mux.HandleFunc("GET /api/users", handlers.GetAllUsers)
	mux.HandleFunc("GET /api/users/{id}", handlers.GetOneUser)
	mux.HandleFunc("POST /api/register", handlers.RegisterUser)
	mux.HandleFunc("POST /api/login", handlers.LoginUser)

	server := Logger(mux)

	log.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", server)
}
