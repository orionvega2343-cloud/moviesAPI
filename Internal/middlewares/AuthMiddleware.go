package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func AuthMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		parseToken := strings.TrimPrefix(tokenStr, "Bearer ")
		token, err := jwt.ParseWithClaims(parseToken, &User{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user", token.Claims.(*User))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
