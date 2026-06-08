package users

import (
	"main/internal/config"

	"github.com/eapache/go-resiliency/retrier"
	"github.com/go-chi/chi"
	"google.golang.org/genproto/googleapis/maps/routes/v1"
)


type Handler struct{
	config *config.Config
	store db.store
}


func NewHandler(cfg *config.Config, store db.store) *Handler{
   return &Handler{
	config: cfg,
	store: store
	
   }
}

func (h *Handler ) Routes() *chi.Mux{
   router := chi.NewRouter
    

   return  router
}