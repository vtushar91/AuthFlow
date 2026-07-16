package services

import "AuthFLow/internal/repository"

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(
	repo *repository.UserRepository,
) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
