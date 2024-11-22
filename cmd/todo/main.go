package main

import (
	"fmt"
	"net/http"
	"todo-api/internal/routes"
	"todo-api/internal/utils"
    "github.com/go-chi/chi/v5"
	"log"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using environment variables")
    }

    utils.GetDB()
	utils.RunMigrations()

	r := chi.NewRouter()
	
    // Mount route groups
    r.Mount("/todo", routes.TodosRoutes())
    r.Mount("/", routes.UsersRoutes())

	fmt.Println("Server is running on port 8080...")

    // Start the server
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	// data.TestConnection()
}
