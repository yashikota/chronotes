package main

import (
	"log"
	"net/http"

	v1 "github.com/yashikota/chronotes/api/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	// Routes
	r.Route("/api/v1", func(r chi.Router) {
		// Debug
		r.HandleFunc("GET /health", v1.HealthHandler)
		r.HandleFunc("GET /provier/github", v1.GithubHandler)
	})

	// SwaggerUI
	swaggerServer := http.StripPrefix("/docs/api", http.FileServer(http.Dir("./docs/api")))
	r.HandleFunc("GET /docs/api/*", func(w http.ResponseWriter, r *http.Request) {
		swaggerServer.ServeHTTP(w, r)
	})

	// Start server
	if err := http.ListenAndServe(":5678", r); err != nil {
		log.Fatal(err)
	}
}
