package main

import (
	"AuthFLow/internal/config"
	"AuthFLow/internal/database"
	"AuthFLow/internal/handlers"
	"AuthFLow/internal/repository"
	"AuthFLow/internal/routes"
	"AuthFLow/internal/services"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.LoadEnv()

	database.ConnectToPostgres()
	defer database.ClosePostgres()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	pool := database.GetDB()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("postgres not ready: %v", err)
	}

	log.Println("postgres connected")

	// ---------------- Repositories ----------------

	userRepo := repository.NewUserRepository(pool)

	// ---------------- Services ----------------

	authService := services.NewAuthService(userRepo)

	// ---------------- Handlers ----------------

	handler := handlers.NewHandler(
		authService,
	)

	router := routes.RegisterRoutes(handler)

	server := &http.Server{
		Addr:              ":7070",
		Handler:           router,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Println("server listening on :7070")

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)

	signal.Notify(
		sigChan,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-sigChan

	log.Println("shutdown initiated")

	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown error: %v", err)
	}

	log.Println("server stopped gracefully")
}
