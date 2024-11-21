package main

import (
	"fmt"
	"net/http"
	"todo-api/handlers/user"
	"todo-api/handlers/todo"
	"todo-api/middleware"
)

func main() {
	// Route Handlers
	// http.HandleFunc("/todos", handlers.CreateTodoHandler)
	http.HandleFunc("/todos/sign-up", user.SignupHandler)
	http.HandleFunc("/todos/login", user.LoginHandler)
	http.HandleFunc("/todos/get-user", user.GetUsersHandler)
	http.HandleFunc("/todos/all", middleware.AuthMiddleware(todo.GetTodosHandler))
	http.HandleFunc("/todos/update", todo.UpdateTodoHandler)
	http.HandleFunc("/todos/delete", todo.DeleteTodoHandler)
	http.HandleFunc("/todos/create", middleware.AuthMiddleware(todo.CreateTodoHandler))

	// Start Server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
