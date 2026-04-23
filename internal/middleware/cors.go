package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		// 🔥 allowed origins (add your frontend URL here)
		allowedOrigins := map[string]bool{
			"http://localhost:3000": true,
			"http://localhost:3001": true,
			"https://peoplespost.vercel.app": true, // 🔥 CHANGE THIS
		}

		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// 🔥 required for cookies
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")

		// 🔥 handle preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}