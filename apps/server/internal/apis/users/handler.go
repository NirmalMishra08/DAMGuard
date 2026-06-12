package users

import (
	"main/internal/config"
	middlewares "main/internal/middleware"
	"main/internal/repository/postgres/sqlc"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	config *config.Config
	store  *sqlc.Store
}

func NewHandler(cfg *config.Config, store *sqlc.Store) *Handler {
	return &Handler{
		config: cfg,
		store:  store,
	}
}

func (h *Handler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/login", h.AuthHandler)

	router.Group(func(r chi.Router) {
		r.Use(middlewares.TokenMiddleware(h.store))

	})

	return router
}
