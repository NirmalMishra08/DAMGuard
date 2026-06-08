package utils

import (
	"main/internal/apis/users"
	"main/internal/config"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func SetupRoutes(router *chi.Mux, cfg *config.Config) {
	// Apply middleware globally
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		// middleware.StripSlashes,
		// middleware.Recoverer,
		// middleware.Heartbeat("/ping"),
		// loggerMiddleware,
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}),
	)

	// Call your route setup functions
	defaultRouter(router, cfg)
}

func defaultRouter(router *chi.Mux, cfg *config.Config) {
	// Health check route
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"message": "pong"})
	})

	// API v1 routes
	router.Route("/api/v1", func(r chi.Router) {
		// User routes
		r.Mount("/users",users.Handler))
	})
}
