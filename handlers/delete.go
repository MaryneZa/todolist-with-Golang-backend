package handlers

import (
	"encoding/json"
	"net/http"
)

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var deleteRequest struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&deleteRequest); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == deleteRequest.ID {
			todos = append(todos[:i], todos[i+1:]...)
			// w.WriteHeader(http.StatusNoContent)
		
			response := map[string]string{
				"message": "To-Do successfully deleted",
			}

			// Set response header for JSON content
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK) // Return 200 OK instead of No Content to send the message
			json.NewEncoder(w).Encode(response) // Encode the response message
			return
		}
	}

	http.Error(w, "To-Do not found", http.StatusNotFound)
}
