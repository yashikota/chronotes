package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/yashikota/chronotes/api/v1/debug"
	"github.com/yashikota/chronotes/api/v1/notes"
	"github.com/yashikota/chronotes/api/v1/upload"
	"github.com/yashikota/chronotes/api/v1/users"
	"github.com/yashikota/chronotes/pkg/db"
	"github.com/yashikota/chronotes/pkg/redis"
	"github.com/yashikota/chronotes/pkg/utils"
)

func main() {
	// Initialize slog
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

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

	// Connect to database
	db.Connect()

	// Initialize Redis
	redis.Connect()
	redis.Initialize()

	// Setup JWT
	utils.SetupPrivateKey()

	r.Route("/api/v1", func(r chi.Router) {
		// Public Routes
		r.HandleFunc("POST /users/register", users.RegisterHandler)
		r.HandleFunc("POST /users/login", users.LoginHandler)

		// Debug Routes
		r.HandleFunc("GET /health", debug.HealthHandler)

		// JWT-protected routes
		r.Group(func(r chi.Router) {
			r.Use(utils.JwtMiddleware)

			// User routes
			r.HandleFunc("POST /users/logout", users.LogoutHandler)
			r.HandleFunc("DELETE /users/me", users.DeleteHandler)
			r.HandleFunc("PUT /users/accounts", users.UpdateAccountsHandler)
			r.HandleFunc("PUT /users/promote", users.PromoteHandler)

			// Notes routes
			r.HandleFunc("GET /notes/note", notes.GetNoteHandler)
			r.HandleFunc("GET /notes/list", notes.GetNoteListHandler)
			r.HandleFunc("GET /notes/summary", notes.GetNoteSummaryHandler)

			// Upload route
			r.HandleFunc("POST /upload/image", upload.UploadHandler)
		})

		// r.Route("/admin", func(r chi.Router) {
		// 	r.Use(utils.AdminMiddleware)

		// 	// Admin-specific routes
		// 	r.HandleFunc("GET /dashboard", admin.DashboardHandler)
		// 	r.HandleFunc("POST /users/manage", admin.ManageUsersHandler)
		// 	r.HandleFunc("DELETE /users/{userID}", admin.DeleteUserHandler)
		// })
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
	slog.Info("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error(err.Error())
	}
}
