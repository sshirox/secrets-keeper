package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/sshirox/secrets-keeper/internal/database"
	"github.com/sshirox/secrets-keeper/internal/handlers"
	"github.com/sshirox/secrets-keeper/internal/middleware"
)

func main() {
	_ = godotenv.Load()
	database.ConnectDatabase()

	r := chi.NewRouter()
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)

	r.Route("/vault", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/", handlers.AddVaultSecret)
		r.Get("/", handlers.GetVaultSecrets)
		r.Delete("/{id}", handlers.DeleteVaultSecret)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Println("The server is running on port: ", port)
	http.ListenAndServe(":"+port, r)
}
