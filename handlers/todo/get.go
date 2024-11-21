package todo

import (
	"encoding/json"
	"net/http"
	"todo-api/data"
)

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	json.NewEncoder(w).Encode(data.Todos)
}
