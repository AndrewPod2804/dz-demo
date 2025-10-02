package middleware

import (
	"4-order-api/configs"
	"4-order-api/pkg/jwt"
	"context"
	"net/http"
	"strings"
)

type key string

const (
	ContextPhonekey key = "ContextPhonekey"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			writeUnauthorized(w)
			return
		}
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		is, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !is {
			writeUnauthorized(w)
			return
		}
		ctx := context.WithValue(r.Context(), ContextPhonekey, data.Phone)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})

}
