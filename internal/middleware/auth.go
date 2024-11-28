package middleware

import (
    "net/http"
    "todo-api/internal/utils"
	"fmt"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("access_token")
        if err != nil {
            http.Error(w, "Access token not provided", http.StatusUnauthorized)
            return
        }

        accessToken := cookie.Value
        userID, err := utils.VerifyToken(accessToken, "access")
        if err != nil {
            http.Error(w, "Invalid or expired access token", http.StatusUnauthorized)
            return
        }

        // Attach user ID to request context
        ctx := context.WithValue(r.Context(), "user_id", userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func AdminMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("user_id").(int)
        
        db := utils.GetDB()
        var role string
        err := db.Get(&role, `SELECT role FROM users WHERE id = ?`, userID)
        if err != nil || role != "admin" {
            http.Error(w, "Unauthorized", http.StatusForbidden)
            return
        }

        next.ServeHTTP(w, r)
    })
}

