package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"word/config"
	"word/internal/service"
	"word/internal/transport/rest/handler"
)

type RestServer struct {
	services *service.Service
}

func New(services *service.Service) *RestServer {
	return &RestServer{
		services: services,
	}
}

func (s *RestServer) Serve() {
	// Context for the server
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Signal to close app
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)

	handlers := handler.New(s.services).Handle()
	server := &http.Server{
		Addr:    config.Port,
		Handler: handlers,
	}

	go func() {
		log.Printf("Starting server on port %s", config.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	select {
	case <-osSignal:
		log.Println("Shutdown signal received. Shutting down...")

		// Timeout to close handlers
		shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 30*time.Second)
		defer shutdownCancel()

		// Stop the server
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Server shutdown failed: %v", err)
		}
	}

}
