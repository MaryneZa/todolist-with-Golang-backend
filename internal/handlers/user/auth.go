package user

import (
    "encoding/json"
    "net/http"
    "todo-api/internal/utils"
    "todo-api/internal/types"
    "database/sql"
    "github.com/go-sql-driver/mysql"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var input types.UserSignUp

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Hash the password
    hashedPassword, err := utils.HashPassword(input.Password)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    // Insert the new user into the database
    db := utils.GetDB()
    query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
    _, err = db.Exec(query, input.Username, input.Email, hashedPassword)
    if err != nil {
        if sqlError, ok := err.(*mysql.MySQLError); ok && sqlError.Number == 1062 {
            http.Error(w, "Email already exists", http.StatusConflict)
        } else {
            http.Error(w, "Error creating user", http.StatusInternalServerError)
        }
        return
    }

    // Respond with success
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Account created successfully",
    })
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var credentials types.UserLoginInput

    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Fetch the user from the database
    db := utils.GetDB()
    var user types.User
    query := `SELECT id, username, email, password FROM users WHERE email = ?`
    err := db.Get(&user, query, credentials.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        } else {
            http.Error(w, "Error fetching user", http.StatusInternalServerError)
        }
        return
    }

    // Verify the password
    if !utils.CheckPasswordHash(credentials.Password, user.Password) {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // Generate a JWT token
    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    // Create a response object excluding the password
    userResponse := types.UserLoginResponse{
        ID:       user.ID,
        Username: user.Username,
        Email:    user.Email,
        Token:    token,
    }

    // Respond with the user and token
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(userResponse)
}
