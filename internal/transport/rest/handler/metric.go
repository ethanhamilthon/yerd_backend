package handler

import (
	"encoding/json"
	"net/http"
	"word/config"
)

func (h *Handler) Visits(w http.ResponseWriter, r *http.Request) {
	Login := r.FormValue("login")
	Pass := r.FormValue("pass")
	if Login != config.AdminLogin || Pass != config.AdminPass {
		http.Error(w, "Wrong login or password", http.StatusUnauthorized)
		return
	}
	visitLogs, err := h.s.Metrics.VisitLogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(visitLogs)
}
