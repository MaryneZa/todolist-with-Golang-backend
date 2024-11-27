package routes

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "todo-api/internal/middleware"
    "todo-api/internal/handlers/user"
)

func UsersRoutes() http.Handler {
    r := chi.NewRouter()

    r.Post("/sign-up", user.SignupHandler)
    r.Post("/login", user.LoginHandler)
    r.Post("/get-access", user.RefreshTokenHandler)
    r.With(middleware.AuthMiddleware).Get("/users", user.GetUsersHandler)

    return r
}
