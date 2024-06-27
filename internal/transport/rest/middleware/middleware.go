package middleware

import (
	"log"
	"net/http"
	"time"
)

func NewLogger() func(func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
		return with(handler, info)
	}
}

func info(next http.HandlerFunc) http.HandlerFunc {
	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("Status: %s %s, in %v", r.Method, r.URL.Path, time.Since(start))
	})
	return hf
}

func with(handler func(http.ResponseWriter, *http.Request), middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	finalHandler := http.HandlerFunc(handler)
	for _, middleware := range middlewares {
		finalHandler = middleware(finalHandler)
	}
	return finalHandler
}
