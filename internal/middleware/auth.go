package middleware

import (
    "net/http"
    "todo-api/internal/utils"
	"fmt"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        userID, err := utils.VerifyJWT(token[7:]) // Strip "Bearer " prefix
        if err != nil {
            http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
            return
        }

        fmt.Printf("Authenticated User ID: %d\n", userID)
        next.ServeHTTP(w, r)
    })
}
