package main

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/yashikota/chronotes/api/v1/auth"
	"github.com/yashikota/chronotes/api/v1/debug"
	"github.com/yashikota/chronotes/api/v1/images"
	"github.com/yashikota/chronotes/api/v1/notes"
	"github.com/yashikota/chronotes/api/v1/users"
	"github.com/yashikota/chronotes/pkg/db"
	"github.com/yashikota/chronotes/pkg/elastic"
	"github.com/yashikota/chronotes/pkg/minio"
	"github.com/yashikota/chronotes/pkg/rabbitmq"
	"github.com/yashikota/chronotes/pkg/redis"
	"github.com/yashikota/chronotes/pkg/utils"
)

func main() {
	// Initialize logger
	logger := utils.Logger()
	slog.SetDefault(logger)

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Initialize
	db.Connect()
	redis.Connect()
	elastic.Connect()
	minio.Connect()
	minio.CreateBucket()
	rabbitmq.Connect()

	// Setup JWT
	utils.LoadPrivateKeyFromEnv()

	r.Route("/api/v1", func(r chi.Router) {
		// Public Routes
		r.HandleFunc("POST /auth/register", auth.RegisterHandler)
		r.HandleFunc("POST /auth/login", auth.LoginHandler)

		// Debug Routes
		r.HandleFunc("GET /health", debug.HealthHandler)
		// r.HandleFunc("GET /debug/users", users.GetUsersListHandler)

		// JWT-protected routes
		r.Group(func(r chi.Router) {
			r.Use(utils.JwtMiddleware)

			// User routes
			r.HandleFunc("POST /auth/logout", auth.LogoutHandler)
			r.HandleFunc("GET /users/me", users.GetUserHandler)
			r.HandleFunc("PUT /users/me", users.UpdateUserHandler)
			r.HandleFunc("DELETE /users/me", users.DeleteUserHandler)
			r.HandleFunc("PUT /users/promote", users.PromoteHandler)

			// Notes routes
			r.HandleFunc("GET /notes", notes.GetNoteHandler)
			r.HandleFunc("PUT /notes", notes.UpdateNoteHandler)
			r.HandleFunc("DELETE /notes", notes.DeleteNoteHandler)
			// r.HandleFunc("POST /notes/search", notes.SearchNoteHandler)
			r.HandleFunc("GET /notes/summary", notes.GetNoteSummaryHandler)
			r.HandleFunc("GET /notes/share", notes.GetSharedNoteHandler)
			r.HandleFunc("POST /notes/share", notes.ShareNoteHandler)
			r.HandleFunc("DELETE /notes/share", notes.UnShareNoteHandler)

			// Images routes
			r.HandleFunc("POST /images", images.UploadImageHandler)
			r.HandleFunc("GET /images", images.GetImageHandler)
			r.HandleFunc("DELETE /images", images.DeleteImageHandler)
		})

		// Admin routes
		r.Route("/admin", func(r chi.Router) {
			r.Use(utils.JwtMiddleware)
			r.Use(utils.AdminMiddleware)

			r.HandleFunc("POST /notes", notes.CreateNoteHandler)
		})
	})

	// Photo Preview
	photoServer := http.StripPrefix("/img/", http.FileServer(http.Dir("./img")))
	r.Get("/img/*", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".jpg") || strings.HasSuffix(r.URL.Path, ".jpeg") || strings.HasSuffix(r.URL.Path, ".png") || strings.HasSuffix(r.URL.Path, ".webp") {
			photoServer.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	// Start server
	slog.Info("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error(err.Error())
	}
}
