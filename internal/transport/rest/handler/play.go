package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) Play(w http.ResponseWriter, r *http.Request) {
	UserID, _, err := CheckAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	countValue := r.FormValue("count")
	count, err := strconv.ParseInt(countValue, 10, 8)
	if err != nil {
		http.Error(w, "count must be number", http.StatusBadRequest)
		return
	}
	langValue := r.FormValue("lang")
	words, err := h.s.Play.GeneratePlay(UserID, int(count), langValue)
	res := map[string]interface{}{
		"count": len(words),
		"words": words,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
