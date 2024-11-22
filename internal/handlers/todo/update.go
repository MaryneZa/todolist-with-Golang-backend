package todo

import (
	// "encoding/json"
	"net/http"
	// "todo-api/internal/types"
)

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// var updatedTodo models.Todo
	// if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
	// 	http.Error(w, "Invalid input", http.StatusBadRequest)
	// 	return
	// }

	// result, err := utils.DB.Exec(
	// 	"UPDATE todos SET title = ?, description = ?, completed = ?, user_id = ? WHERE id = ?",
	// 	input.Title, input.Description, input.Completed, input.UserID, input.ID,
	// )
	// if err != nil {
	// 	http.Error(w, "Failed to update todo", http.StatusInternalServerError)
	// 	return
	// }
	
	// rowsAffected, _ := result.RowsAffected()
	// if rowsAffected == 0 {
	// 	http.Error(w, "No todo found with the given ID", http.StatusNotFound)
	// 	return
	// }
	

	w.WriteHeader(http.StatusOK)
}
