package handler

import (
	"encoding/json"
	"net/http"
)

type askBody struct {
	ID     string `json:"id"`
	Oslang string `json:"oslang"`
	Tolang string `json:"tolang"`
	Word   string `json:"word"`
}

type FlushWriter struct {
	flush http.Flusher
	wHttp http.ResponseWriter
}

func NewWriter(flush http.Flusher, wHttp http.ResponseWriter) *FlushWriter {
	return &FlushWriter{
		flush: flush,
		wHttp: wHttp,
	}
}

func (w *FlushWriter) Write(p []byte) (int, error) {
	n, err := w.wHttp.Write(p)
	w.flush.Flush()
	return n, err
}

func (h *Handler) Ask(w http.ResponseWriter, r *http.Request) {
	UserID, _, err := CheckAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get data from body
	var askparams askBody
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err = decoder.Decode(&askparams)
	if err != nil {
		http.Error(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	//Set headers for streaming
	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	//Create flusher
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Error on creating flusher", http.StatusBadRequest)
		return
	}

	//Create Writer function
	Writer := NewWriter(flusher, w)

	err = h.s.Ask.GenerateWord(askparams.ID, UserID, askparams.Oslang, askparams.Tolang, askparams.Word, Writer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}
