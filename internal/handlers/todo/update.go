package todo

import (
	"encoding/json"
	"net/http"
	"todo-api/internal/types"
	"todo-api/internal/utils"
)


func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse and decode the JSON body into the Todo struct
	var updatedTodo types.UpdateTodo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Ensure required fields are present
	// Validate the input struct
	if err := validate.Struct(updatedTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the todo in the database
	db := utils.GetDB()
	query := `UPDATE todos SET title = ?, description = ?, completed = ?, user_id = ? WHERE id = ?`
	result, err := db.Exec(query, updatedTodo.Title, updatedTodo.Description, updatedTodo.Completed, updatedTodo.UserID, updatedTodo.ID)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	// Check if the todo was updated
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to confirm update", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No todo found with the given ID", http.StatusNotFound)
		return
	}

	// Respond with success
	response := map[string]string{
		"message": "To-Do successfully updated",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
