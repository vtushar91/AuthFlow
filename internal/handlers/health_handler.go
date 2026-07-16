package handlers

import (
	"AuthFLow/internal/database"
	"encoding/json"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Timestamp time.Time `json:"timestamp"`
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := database.GetDB().Ping(r.Context()); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)

		json.NewEncoder(w).Encode(HealthResponse{
			Status:    "down",
			Service:   "AuthFlow",
			Timestamp: time.Now().UTC(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(HealthResponse{
		Status:    "ok",
		Service:   "AuthFlow",
		Timestamp: time.Now().UTC(),
	})
}
