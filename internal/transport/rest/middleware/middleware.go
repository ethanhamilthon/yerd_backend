package middleware

import (
	"log"
	"net/http"
	"time"
)

type Storage interface {
	Visit(int, string, string, string) error
}

func NewLogger(s Storage) func(func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
		return with(s, handler, info)
	}
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Status() int {
	if rw.status == 0 {
		return http.StatusOK
	}
	return rw.status
}

func info(next http.HandlerFunc, s Storage) http.HandlerFunc {
	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		wrapped := &responseWriter{ResponseWriter: w}
		next.ServeHTTP(wrapped, r)

		log.Printf("Status:%v Path: %s %s, in %v", wrapped.status, r.Method, r.URL.Path, time.Since(start))
		s.Visit(wrapped.status, r.URL.Path, r.Method, time.Since(start).String())
	})
	return hf
}

func with(s Storage, handler func(http.ResponseWriter, *http.Request), middlewares ...func(http.HandlerFunc, Storage) http.HandlerFunc) http.HandlerFunc {
	finalHandler := http.HandlerFunc(handler)
	for _, middleware := range middlewares {
		finalHandler = middleware(finalHandler, s)
	}
	return finalHandler
}
