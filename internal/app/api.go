package app

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"dev/internal/app/handler"
)

func NewChiHandeler() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Use(middleware.DefaultLogger)
		r.Use(middleware.Timeout(30 * time.Second))
		r.Use()

		r.Post("/user", handler.UserCreate)
		r.Post("/team", handler.TeamCreate)
		r.Post("/hub", handler.HubCreate)

		r.Get("/search", handler.Search)
	})

	return r
}
