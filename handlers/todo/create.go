package todo

import (
	"encoding/json"
	"net/http"
	"todo-api/data"
	"todo-api/models"
)

// todos := data.Todos
// var nextID = 3

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

	nextID := len(data.Todos) + 1
	todo.ID = nextID

	data.Todos = append(data.Todos, todo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
