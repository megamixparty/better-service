package middlewares

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		matched, _ := regexp.MatchString(`auth.*`, r.URL.Path)
		if !matched {
			res := GetRedis(r.Context()).Get(r.Header.Get("token"))
			if res.Err() != nil {
				w.WriteHeader(http.StatusUnauthorized)
				response := ErrorResponse{"Unauthorize. Please login!"}
				b, _ := json.Marshal(response)
				w.Write(b)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
