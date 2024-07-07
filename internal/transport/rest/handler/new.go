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
	return &Handler{s: s}
}

func (h *Handler) Handle() *http.ServeMux {
	//Define the main route
	handlers := http.NewServeMux()
	middl := middleware.NewLogger(h.s.Metrics)

	//Define the api routes
	apiv1 := http.NewServeMux()
	apiv1.HandleFunc("POST /word", middl(h.CreateWord))
	apiv1.HandleFunc("PATCH /word", middl(h.UpdateWord))
	apiv1.HandleFunc("DELETE /word/{id}", middl(h.DeleteWord))
	apiv1.HandleFunc("/word/{id}", middl(h.Word))
	apiv1.HandleFunc("/word", middl(h.Words))
	apiv1.HandleFunc("PATCH /onboard", middl(h.OnboardUpdate))
	apiv1.HandleFunc("/play", middl(h.Play))
	apiv1.HandleFunc("POST /ask", h.Ask)
	apiv1.HandleFunc("/me", middl(h.Me))
	apiv1.HandleFunc("/metrics/visits", middl(h.Visits))
	handlers.Handle("/api/v1/", http.StripPrefix("/api/v1", apiv1))

	//Define oauth2 routes
	oauth2_google := http.NewServeMux()
	oauth2_google.HandleFunc("/login", middl(h.GoogleLoginURL))
	oauth2_google.HandleFunc("/callback", middl(h.GoogleCallback))

	handlers.Handle("/oauth/google/", http.StripPrefix("/oauth/google", oauth2_google))

	return handlers
}
