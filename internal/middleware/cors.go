package middleware

import "net/http"

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// These are permissive by default but should eventually be more specific
		// There are a lot more headers you can add here for security reasons.
		// Check https://www.securityheaders.com for other examples
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origins", "*")
		next.ServeHTTP(w, r)
	})
}
