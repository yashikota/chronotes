package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/stretchr/testify/assert"

	"github.com/yashikota/chronotes/api/v1/debug"
	"github.com/yashikota/chronotes/api/v1/users"
)

func TestRoutes(t *testing.T) {
	r := setupRouter()

	t.Run("GET /api/v1/health", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/health", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	// FIXME: Below routes are not implemented yet
	// t.Run("GET /api/v1/provider/github", func(t *testing.T) {
	// 	req, _ := http.NewRequest("GET", "/api/v1/provider/github", nil)
	// 	resp := httptest.NewRecorder()
	// 	r.ServeHTTP(resp, req)

	// 	assert.Equal(t, http.StatusOK, resp.Code)
	// })

	// t.Run("GET /api/v1/provider/discord", func(t *testing.T) {
	// 	req, _ := http.NewRequest("GET", "/api/v1/provider/discord", nil)
	// 	resp := httptest.NewRecorder()
	// 	r.ServeHTTP(resp, req)

	// 	assert.Equal(t, http.StatusOK, resp.Code)
	// })

	t.Run("GET /docs/api/*", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/docs/api/index.html", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusMovedPermanently, resp.Code)
	})
}

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
		r.HandleFunc("POST /users/register", users.RegisterHandler)
		r.HandleFunc("POST /users/login", users.LoginHandler)

		// Debug
		r.HandleFunc("GET /health", debug.HealthHandler)
	})
	// SwaggerUI
	swaggerServer := http.StripPrefix("/docs/api", http.FileServer(http.Dir("/app/docs/api")))
	r.HandleFunc("GET /docs/api/*", func(w http.ResponseWriter, r *http.Request) {
		swaggerServer.ServeHTTP(w, r)
	})

	return r
}
