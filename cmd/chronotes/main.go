package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/yashikota/chronotes/api/v1/debug"
	"github.com/yashikota/chronotes/api/v1/users"
	"github.com/yashikota/chronotes/pkg/db"
	"github.com/yashikota/chronotes/pkg/utils"
)

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Connect to database
	db.Connect()

	// Setup JWT
	utils.SetupPrivateKey()

	// Routes
	r.Route("/api/v1", func(r chi.Router) {
		r.HandleFunc("POST /users/register", users.RegisterHandler)
		r.HandleFunc("POST /users/login", users.LoginHandler)

		// Debug
		r.HandleFunc("GET /health", debug.HealthHandler)
	})

	// Routes with JWT middleware
	// r.Route("/api/v1", func(r chi.Router) {
	// r.Use(v1.JwtMiddleware)

	// User
	// r.HandleFunc("POST /users/logout", v1.LogoutHandler)
	// r.HandleFunc("DELETE /users/{user_id}", v1.DeleteUserHandler)

	// Providers
	// r.HandleFunc("GET /provider/github", provider.GithubHandler)
	// r.HandleFunc("GET /provider/discord", provider.DiscordHandler)
	// })

	// Start server
	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
