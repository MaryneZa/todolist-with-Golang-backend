package todo

import (
	"encoding/json"
	"net/http"
	"todo-api/internal/types"
	"todo-api/internal/utils"
)

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Fetch todos from the database
	db := utils.GetDB()
	var todos []types.Todo
	query := `SELECT id, title, description, completed, user_id FROM todos`
	err := db.Select(&todos, query)
	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}

	// Respond with the list of todos
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
