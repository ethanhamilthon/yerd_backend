package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"word/config"
)

func (h *Handler) GoogleLoginURL(w http.ResponseWriter, r *http.Request) {
	url := h.s.User.GoogleLoginURL()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	code := r.FormValue("code")
	UserID, Email, err := h.s.User.GoogleCallback(state, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := CreateJWT(UserID, Email)
	if err != nil {
		http.Error(w, "Token not created", http.StatusBadRequest)
		return
	}
	cookie := createCookie(time.Now().Add(24*60*time.Hour), "Authorization", token)
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, config.RedirectUser, http.StatusTemporaryRedirect)

}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	UserID, Email, err := CheckAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, languages, err := h.s.User.User(UserID, Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := map[string]interface{}{
		"user":      user,
		"languages": languages,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

type OnboardBody struct {
	OsLanguage      string   `json:"os_language"`
	TargetLanguages []string `json:"target_languages"`
}

func (h *Handler) OnboardUpdate(w http.ResponseWriter, r *http.Request) {
	UserID, _, err := CheckAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var body OnboardBody
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err = decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	err = h.s.User.UpdateLanguages(body.OsLanguage, body.TargetLanguages, UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
