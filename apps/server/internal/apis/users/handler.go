package users

import (
	"main/internal/config"
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

	return router
}
