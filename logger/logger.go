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
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}
