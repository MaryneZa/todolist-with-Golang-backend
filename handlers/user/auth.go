package user

import (
    "encoding/json"
    "net/http"
    "todo-api/data"
    "todo-api/models"
    "todo-api/utils"
    "fmt"
    "strings"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Hash the password
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }
    user.Password = hashedPassword

    // Add user to data store
    user.ID = len(data.Users) + 1
    data.Users = append(data.Users, user)
    fmt.Println(data.Users) 
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var credentials models.User
    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Verify user
    for _, user := range data.Users {
        if user.Username == strings.TrimSpace(credentials.Username) {
            if utils.CheckPasswordHash(strings.TrimSpace(credentials.Password), user.Password) {
                token, err := utils.GenerateJWT(user.ID)
                if err != nil {
                    http.Error(w, "Error generating token", http.StatusInternalServerError)
                    return
                }

                json.NewEncoder(w).Encode(map[string]string{"token": token})
                return
            }
        }
    }

    http.Error(w, "Invalid username or password", http.StatusUnauthorized)
}
