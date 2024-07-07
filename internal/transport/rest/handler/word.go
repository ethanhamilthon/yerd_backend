package handler

import (
	"encoding/json"
	"net/http"
	"word/internal/entities"
)

func (h *Handler) Words(w http.ResponseWriter, r *http.Request) {
	UserID, _, err := CheckAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user_words, err := h.s.Word.UserWords(UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user_words)
}

func (h *Handler) CreateWord(w http.ResponseWriter, r *http.Request) {
	UserID, _, err := CheckAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var word entities.WordBasic
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err = decoder.Decode(&word)
	if err != nil {
		http.Error(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	err = h.s.Word.CreateManualWord(word, UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(word)
}

func (h *Handler) Word(w http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("id")

	word, err := h.s.Word.Word(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(word)
}

func (h *Handler) DeleteWord(w http.ResponseWriter, r *http.Request) {
	UserID, _, err := CheckAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ID := r.PathValue("id")

	err = h.s.Word.DeleteWord(ID, UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) UpdateWord(w http.ResponseWriter, r *http.Request) {
	UserID, _, err := CheckAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var word entities.WordBasic
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err = decoder.Decode(&word)
	if err != nil {
		http.Error(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	err = h.s.Word.UpdateWord(word.ID, word.Title, word.Description, UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
