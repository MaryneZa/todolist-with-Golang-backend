package data

import "todo-api/models"

// Todos is the in-memory storage for to-do items.
var Todos = []models.Todo{
    // Example data
    {ID: 1, Title: "Learn Go", Description: "Start learning Go basics.", Completed: false, UserID: 1},
    {ID: 2, Title: "Build a To-Do App", Description: "Develop a simple to-do application.", Completed: false, UserID: 2},
}
