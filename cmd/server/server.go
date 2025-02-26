package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/sshirox/secrets-keeper/internal/database"
	"github.com/sshirox/secrets-keeper/internal/handlers"
	auth "github.com/sshirox/secrets-keeper/internal/middleware"
)

func main() {
	_ = godotenv.Load()
	database.ConnectDatabase()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)

	r.Route("/vault", func(r chi.Router) {
		r.Use(auth.AuthMiddleware)

		r.Post("/", handlers.AddVaultEntry)
		r.Get("/", handlers.GetVaultEntries)
		r.Delete("/{id}", handlers.DeleteVaultEntry)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("The server is running on port: ", port)
	http.ListenAndServe(":"+port, r)
}
