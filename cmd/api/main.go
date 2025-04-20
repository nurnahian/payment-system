package main

import (
	"log"
	"net/http"
	"os"

	"payment-system/internal/common/db"
	"payment-system/internal/middleware"
	"payment-system/internal/user"

	"github.com/gorilla/mux"
	//"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db.Connect()

	repo := user.NewRepository()
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	r := mux.NewRouter()

	// Apply rate limiting middleware globally
	r.Use(middleware.RateLimitMiddleware)

	// Apply JWT middleware
	r.Use(middleware.JWTMiddleware)

	// Public routes
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/register", handler.Register).Methods("POST")

	// Protected routes
	r.HandleFunc("/users", handler.GetAll).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
