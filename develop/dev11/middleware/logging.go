package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("\nMethod: %s\n URI: %s\n Time: %s\n", r.Method, r.RequestURI, time.Since(timeStart))
	})
}
