package todo

import (
	"encoding/json"
	"net/http"
	"todo-api/models"
	"todo-api/data"
)

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updatedTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, todo := range data.Todos {
		if todo.ID == updatedTodo.ID {
			data.Todos[i] = updatedTodo
			json.NewEncoder(w).Encode(updatedTodo)
			return
		}
	}

	http.Error(w, "To-Do not found", http.StatusNotFound)
}
