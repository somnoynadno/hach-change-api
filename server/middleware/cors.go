package middleware

import (
	u "hack-change-api/muxutil"
	"net/http"
	"strings"
)

// CORS middleware adds necessary access-control headers.
var CORS = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			w.Header().Add("Content-Type", "application/json")
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Add("X-Frame-Options", "DENY")

		if r.Method == http.MethodOptions {
			u.HandleOptions(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
