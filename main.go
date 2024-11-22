package main

import (
	"fmt"
	"net/http"
	"todo-api/routes"
    "github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

    // Mount route groups
    r.Mount("/todo", routes.TodosRoutes())
    r.Mount("/", routes.UsersRoutes())

	fmt.Println("Server is running on port 8080...")

    // Start the server
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
