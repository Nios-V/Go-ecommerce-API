package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Nios-V/Go-ecommerce-API/internal/response"
	"github.com/Nios-V/Go-ecommerce-API/internal/security"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(w, http.StatusUnauthorized, "Missing auth token")
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := security.ValidateJWT(token)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "email", claims.Email)
		ctx = context.WithValue(ctx, "roles", claims.Roles)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
