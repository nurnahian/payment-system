package main

import (
	"log"
	"net/http"
	"os"

	"payment-system/internal/common/db"
	"payment-system/internal/user"
	"payment-system/internal/middleware"

	"github.com/gorilla/mux"
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

	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/users", handler.GetAll).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
