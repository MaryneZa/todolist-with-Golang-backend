package todo

import (
	"encoding/json"
	"net/http"
	"todo-api/internal/types"
	"todo-api/internal/utils"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse and decode the JSON body into CreateTodoInput struct
	var input types.CreateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate the input struct
	if err := validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the new todo into the database
	db := utils.GetDB()
	query := `INSERT INTO todos (title, description, completed, user_id) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(query, input.Title, input.Description, input.Completed, input.UserID)
	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	// Retrieve the auto-generated ID of the new todo
	todoID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve new todo ID", http.StatusInternalServerError)
		return
	}

	// Respond with the created todo ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Todo created successfully",
		"id":      todoID,
	})
}
