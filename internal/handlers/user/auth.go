package user

import (
    "encoding/json"
    "net/http"
    "todo-api/internal/utils"
    "todo-api/internal/types"
    "database/sql"
    "github.com/go-sql-driver/mysql"
	"github.com/go-playground/validator/v10"
    "time"
    "log"
)

var validate = validator.New()

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

    // Validate the input struct
	if err := validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

    // Validate the input struct
	if err := validate.Struct(credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
    accessToken, err := utils.GenerateAccessTokenJWT(user.ID)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }
    refreshToken, err := utils.GenerateRefreshToken(user.ID)
    if err != nil {
        http.Error(w, "Error generating refresh token", http.StatusInternalServerError)
        return
    }

    // Store refresh token in database
    query = `INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES (?, ?, ?)`
    _, err = db.Exec(query, user.ID, refreshToken, utils.RefreshTokenExpiration())
    if err != nil {
        http.Error(w, "Error saving refresh token", http.StatusInternalServerError)
        return
    }

     // Set cookies for tokens
    http.SetCookie(w, &http.Cookie{
        Name:     "access_token",
        Value:    accessToken,
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteStrictMode,
        Expires:  time.Now().Add(15 * time.Minute),
    })

    http.SetCookie(w, &http.Cookie{
        Name:     "refresh_token",
        Value:    refreshToken,
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteStrictMode,
        Expires:  utils.RefreshTokenExpiration(),
    })

    // Respond with success
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Login successful",
    })
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Extract refresh token from cookie
    cookie, err := r.Cookie("refresh_token")
    if err != nil {
        http.Error(w, "Refresh token not provided", http.StatusUnauthorized)
        return
    }
    refreshToken := cookie.Value
    log.Println(refreshToken)

    // Validate refresh token
    db := utils.GetDB()

    var tokenData struct {
        UserID    int       `db:"user_id"`    // Maps the "user_id" column
        ExpiresAt time.Time `db:"expires_at"` // Maps the "expires_at" column
    }

    query := `SELECT user_id, expires_at FROM refresh_tokens WHERE token = ?`
    err = db.Get(&tokenData, query, refreshToken)
    log.Println(err)
    log.Println(tokenData)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
        } else {
            http.Error(w, "Error validating refresh token", http.StatusInternalServerError)
        }
        return
    }

    if time.Now().After(tokenData.ExpiresAt) {
        http.Error(w, "Refresh token expired", http.StatusUnauthorized)
        return
    }

    // Generate new access token
    accessToken, err := utils.GenerateAccessTokenJWT(tokenData.UserID)
    if err != nil {
        http.Error(w, "Error generating access token", http.StatusInternalServerError)
        return
    }

    // Optionally, rotate the refresh token
    newRefreshToken, err := utils.GenerateRefreshToken(tokenData.UserID)
    if err != nil {
        http.Error(w, "Error generating new refresh token", http.StatusInternalServerError)
        return
    }

    updateQuery := `UPDATE refresh_tokens SET token = ?, expires_at = ? WHERE token = ?`
    _, err = db.Exec(updateQuery, newRefreshToken, utils.RefreshTokenExpiration(), refreshToken)
    if err != nil {
        http.Error(w, "Error updating refresh token", http.StatusInternalServerError)
        return
    }

    // Set cookies for the new tokens
    http.SetCookie(w, &http.Cookie{
        Name:     "access_token",
        Value:    accessToken,
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteStrictMode,
        Expires:  time.Now().Add(15 * time.Minute),
    })

    http.SetCookie(w, &http.Cookie{
        Name:     "refresh_token",
        Value:    newRefreshToken,
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteStrictMode,
        Expires:  utils.RefreshTokenExpiration(),
    })

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Access token refreshed",
    })
}