package handlers

import (
	"encoding/json"
	"net/http"
	"todo-api/models"
)

var todos []models.Todo
var nextID = 1

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	todo.ID = nextID
	nextID++
	todos = append(todos, todo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
