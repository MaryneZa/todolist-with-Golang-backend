package routes

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "todo-api/middleware"
    "todo-api/handlers/todo"
)

func TodosRoutes() http.Handler {
    r := chi.NewRouter()

    r.Put("/update", todo.UpdateTodoHandler)
    r.Delete("/delete", todo.DeleteTodoHandler)
    r.With(middleware.AuthMiddleware).Get("/all", todo.GetTodosHandler)
    r.With(middleware.AuthMiddleware).Post("/create", todo.CreateTodoHandler)
    return r
}
