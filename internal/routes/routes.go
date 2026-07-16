package routes

import (
	"AuthFLow/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(
	h *handlers.Handler,
) *chi.Mux {

	r := chi.NewRouter()

	//health check
	r.Get("/send-otp", h.HealthCheck)

	r.Route("/auth", func(r chi.Router) {

		r.Post("/send-otp", h.Register)

	})
	return r
}
