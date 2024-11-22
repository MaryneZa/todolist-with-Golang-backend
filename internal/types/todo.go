package types

type Todo struct {
    ID        int    `db:"id" json:"id"`
    Title     string `db:"title" json:"title"`
    Description string `db:"description" json:"description"`
    Completed bool   `db:"completed" json:"completed"`
    UserID    int    `db:"user_id" json:"user_id"`
}

type CreateTodoInput struct {
    Title       string `json:"title" validate:"required"`        // Required
    Description string `json:"description"`                     // Optional
    Completed   bool   `json:"completed"`                       // Defaults to false
    UserID      int    `json:"user_id" validate:"required"`      // Required
}

