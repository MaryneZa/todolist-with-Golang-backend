package routes

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "todo-api/middleware"
    "todo-api/handlers/user"
)

func UsersRoutes() http.Handler {
    r := chi.NewRouter()

    r.Post("/sign-up", user.SignupHandler)
    r.Post("/login", user.LoginHandler)
    r.With(middleware.AuthMiddleware).Get("/users", user.GetUsersHandler)

    return r
}