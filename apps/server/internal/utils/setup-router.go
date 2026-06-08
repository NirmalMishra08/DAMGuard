package utils

import (
	"main/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/cors"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config) {
	// Apply middleware globally
	router.Use(
		middleware.SetContentType(), // Gin middleware for JSON
		middleware.StripSlashes,
		middleware.Recoverer,
		middleware.Heartbeat("/ping"),
		loggerMiddleware,
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}),
	)

	// Then call your route setup functions
	defaultRouter(router, cfg) // or however you want to organize routes
}
