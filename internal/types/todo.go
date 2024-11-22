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

type DeleteRequest struct {
    ID int `json:"id"`
}

type UpdateTodo struct {
    ID        int    `db:"id" json:"id" validate:"required"`
    Title     string `db:"title" json:"title" validate:"required"`
    Description string `db:"description" json:"description" validate:"required"`
    Completed bool   `db:"completed" json:"completed" validate:"required"`
    UserID    int    `db:"user_id" json:"user_id" validate:"required"`
}

