package data

import "todo-api/models"

// Users is the in-memory storage for user data.
var Users = []models.User{
    // Example data
	{ID: 1, Username: "maryne", Password: "$2a$10$BGLzk8FtoNQ0jyvZhpY8puqq8GU7gRnyWmE7r5Q798/52HyeLlKkW"}, // 123
	{ID: 3, Username: "mare", Password: "$2a$10$asPwa4Y06B1fcnnzPk2kKeMqitj9tTWeb1W1pLBw.mO8rCwucQcwa"}, // 123
}
