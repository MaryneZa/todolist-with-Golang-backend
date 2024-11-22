package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"todo-api/internal/types"
	"todo-api/internal/utils"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Fetch users from the database
	db := utils.GetDB()
	var users []types.User
	query := `SELECT id, username, email FROM users`
	err := db.Select(&users, query)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No users found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching users", http.StatusInternalServerError)
		}
		return
	}

	// Respond with the list of users
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
