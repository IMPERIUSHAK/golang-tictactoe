package middleware

import (
	"backend/utils"
	"log"
	"net/http"
	"time"
)

func Logs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		log.Printf("[%s] %s %s -%s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			time.Since(start))

		log.Printf(" token %s", string(utils.Jwtsecret))
		next.ServeHTTP(w, r)
	})
}
