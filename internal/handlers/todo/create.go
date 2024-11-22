package todo

import (
	"encoding/json"
	"net/http"
	// "todo-api/internal/utils"
	// "todo-api/internal/types"
)

// todos := data.Todos
// var nextID = 3

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// var todo models.Todo
	// if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
	// 	http.Error(w, "Invalid input", http.StatusBadRequest)
	// 	return
	// }

	// db := utils.GetDB()
    // _, err := db.NamedExec(`INSERT INTO todos (title, completed, user_id) VALUES (:title, :completed, :user_id)`, todo)
    // if err != nil {
    //     http.Error(w, "Failed to create todo", http.StatusInternalServerError)
    //     return
    // }

	// w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Todo created"})
}
