package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/yashikota/chronotes/api/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	r := setupRouter()

	t.Run("GET /api/v1/health", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/health", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("GET /api/v1/provier/github", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/provier/github", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("GET /api/v1/provier/discord", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/provier/discord", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("GET /docs/api/*", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/docs/api/index.html", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code) // Assuming docs directory does not exist
	})
}

// setupRouter initializes the router with middleware and routes for testing.
func setupRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware
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
		r.HandleFunc("GET /health", v1.HealthHandler)
		r.HandleFunc("GET /provier/github", v1.GithubHandler)
		r.HandleFunc("GET /provier/discord", v1.DiscordHandler)
	})

	// SwaggerUI (Mock)
	swaggerServer := http.StripPrefix("/docs/api", http.FileServer(http.Dir("/app/docs/api")))
	r.HandleFunc("GET /docs/api/*", func(w http.ResponseWriter, r *http.Request) {
		swaggerServer.ServeHTTP(w, r)
	})

	return r
}
