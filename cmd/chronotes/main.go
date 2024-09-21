package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/yashikota/chronotes/api/v1/debug"
	"github.com/yashikota/chronotes/api/v1/upload"
	"github.com/yashikota/chronotes/api/v1/users"
	"github.com/yashikota/chronotes/pkg/db"
	"github.com/yashikota/chronotes/pkg/redis"
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

	// Initialize Redis
	redis.Connect()
	redis.Initialize()

	// Setup JWT
	utils.SetupPrivateKey()

	// Public Routes
	r.HandleFunc("POST /api/v1/users/register", users.RegisterHandler)
	r.HandleFunc("POST /api/v1/users/login", users.LoginHandler)

	// Debug
	r.HandleFunc("GET /api/v1/health", debug.HealthHandler)
	r.HandleFunc("GET /api/v1/fake", debug.FakeHandler)

	// Routes with JWT middleware
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(utils.JwtMiddleware)

		// User
		r.HandleFunc("POST /users/logout", users.LogoutHandler)
		r.HandleFunc("DELETE /users/{id}", users.DeleteHandler)

		// Upload
		r.HandleFunc("POST /upload/image", upload.UploadHandler)

		// Providers
		// r.HandleFunc("GET /provider/github", provider.GithubHandler)
		// r.HandleFunc("GET /provider/discord", provider.DiscordHandler)
	})

	// Photo Preview
	photoServer := http.StripPrefix("/img/", http.FileServer(http.Dir("./img")))
	r.Get("/img/*", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".jpg") || strings.HasSuffix(r.URL.Path, ".jpeg") || strings.HasSuffix(r.URL.Path, ".png") {
			photoServer.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	// Start server
	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
