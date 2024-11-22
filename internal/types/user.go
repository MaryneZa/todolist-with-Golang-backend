package types

type User struct {
    ID       int    `db:"id" json:"id"`
    Username string `db:"username" json:"username"`
    Password string `db:"password" json:"password"` // Hashed
    Email    string `db:"email" json:"email"`
}

type UserSignUp struct {
    Username string `json:"username" validate:"required"`
    Email    string `json:"email" validate:"required"`
    Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Token    string `json:"token"`
}

type UserLoginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
