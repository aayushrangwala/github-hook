package util

import (
	"log"
	"net/http"
	"time"
)

// Logger function is the utility helper function which takes the hendler and its name as input for proper logging format
func Logger(inner http.Handler, name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	}
}
