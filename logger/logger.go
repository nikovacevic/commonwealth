package logger

import (
	"log"
	"net/http"
	"time"
)

// LogRequest wraps a HandlerFunc with logging to stdout
func LogRequest(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		hf(w, r)
		log.Printf(
			"%-6s\t%-30s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}

// LogError wraps an erroneous request's HandlerFunc with logging to stdout
func LogError(code string, hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		hf(w, r)
		log.Printf(
			"%-6s\t%-30s\t%s",
			code,
			r.RequestURI,
			time.Since(start),
		)
	}
}
