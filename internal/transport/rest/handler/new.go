package handler

import (
	"net/http"
	"word/internal/service"
	"word/internal/transport/rest/middleware"
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
	logger := middleware.NewLogger()

	//Define the api routes
	apiv1 := http.NewServeMux()
	apiv1.HandleFunc("POST /word", logger(h.CreateWord))
	apiv1.HandleFunc("PATCH /word", logger(h.UpdateWord))
	apiv1.HandleFunc("DELETE /word/{id}", logger(h.DeleteWord))
	apiv1.HandleFunc("/word/{id}", logger(h.Word))
	apiv1.HandleFunc("/word", logger(h.Words))
	apiv1.HandleFunc("PATCH /onboard", logger(h.OnboardUpdate))
	apiv1.HandleFunc("/play", logger(h.Play))
	apiv1.HandleFunc("POST /ask", h.Ask)
	apiv1.HandleFunc("/me", logger(h.Me))
	handlers.Handle("/api/v1/", http.StripPrefix("/api/v1", apiv1))

	//Define oauth2 routes
	oauth2_google := http.NewServeMux()
	oauth2_google.HandleFunc("/login", logger(h.GoogleLoginURL))
	oauth2_google.HandleFunc("/callback", logger(h.GoogleCallback))

	handlers.Handle("/oauth/google/", http.StripPrefix("/oauth/google", oauth2_google))

	return handlers
}
