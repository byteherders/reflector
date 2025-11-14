package server

import (
	"log"
	"net/http"
	"time"
)

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s from %s (%s) in %s", r.Method, r.URL.String(), clientIP(r), r.RemoteAddr, time.Since(start))
	})
}
