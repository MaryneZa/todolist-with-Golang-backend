package data

import "todo-api/models"

// Users is the in-memory storage for user data.
var Users = []models.User{
    // Example data
	{ID: 1, Username: "maryne", Password: "$2a$10$20/54Htif5xuSBVhrzOvL..GXdrQChChr9iCa4Y4RcCsUHmZW8dwq"}, // 123
	{ID: 3, Username: "mare", Password: "$2a$10$4ywHVbiS0L0Uu.THBMM/3epx9MjFD4M2Pzsz8lsnvbj5bgWdceIhq"}, // 123
}
