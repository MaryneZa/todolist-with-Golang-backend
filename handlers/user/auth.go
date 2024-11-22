package user

import (
    "encoding/json"
    "fmt"
    "net/http"
    "todo-api/utils"
    "todo-api/data"
    "todo-api/models"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var user struct {
        ID       int    `json:"id"`
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Simulate storing the user (you should use a database in a real app)
    fmt.Printf("User created: %+v\n", user)


    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    hashed, err := utils.HashPassword(user.Password)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }


    json.NewEncoder(w).Encode(map[string]string{
        "token": token, 
        "hashed password": hashed,
    })
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse the login credentials
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Simulate fetching user data from in-memory storage or database
    var storedUser *models.User
    for _, user := range data.Users { // Assuming `data.Users` is a slice of User
        if user.Username == credentials.Username {
            storedUser = &user
            break
        }
    }

    // Handle case where user is not found
    if storedUser == nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    // Verify the password
    if !utils.CheckPasswordHash(credentials.Password, storedUser.Password) {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    // Generate a JWT token
    token, err := utils.GenerateJWT(storedUser.ID)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    // Respond with the token
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
