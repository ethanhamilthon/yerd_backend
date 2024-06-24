package handler

import (
	"net/http"
	"word/internal/service"
)

type Handler struct {
	s *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Handle() *http.ServeMux {
	//Define the main route
	handlers := http.NewServeMux()

	//Define the api routes
	api := http.NewServeMux()
	api.HandleFunc("/ask", h.Ask)

	return handlers
}
