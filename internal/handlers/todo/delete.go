package todo

import (
	"encoding/json"
	"net/http"
	"todo-api/internal/utils"
	"todo-api/internal/types"

)

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body to get the ID of the todo to delete
	var deleteRequest types.DeleteRequest
	
	if err := json.NewDecoder(r.Body).Decode(&deleteRequest); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Ensure the ID is provided and valid
	if deleteRequest.ID <= 0 {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	// Connect to the database and attempt to delete the todo
	db := utils.GetDB()
	result, err := db.Exec(`DELETE FROM todos WHERE id = ?`, deleteRequest.ID)
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	// Check if any rows were affected (i.e., if the todo existed)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to confirm deletion", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "To-Do not found", http.StatusNotFound)
		return
	}

	// Respond with a success message
	response := map[string]string{
		"message": "To-Do successfully deleted",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
