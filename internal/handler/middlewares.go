package handler

import "net/http"

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	}
}
