package handlers

import "AuthFLow/internal/services"

type Handler struct {
	AuthService *services.AuthService
}

func NewHandler(
	AuthService *services.AuthService,
) *Handler {
	return &Handler{
		AuthService: AuthService,
	}
}
