package main

import (
	"fmt"
	"net/http"
	"todo-api/handlers"
)

func main() {
	// Route Handlers
	http.HandleFunc("/todos", handlers.CreateTodoHandler)
	http.HandleFunc("/todos/all", handlers.GetTodosHandler)
	http.HandleFunc("/todos/update", handlers.UpdateTodoHandler)
	http.HandleFunc("/todos/delete", handlers.DeleteTodoHandler)

	// Start Server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
