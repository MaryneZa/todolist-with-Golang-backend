package utils

import (
    "log"
)

func RunMigrations() {
    db := GetDB()
    queries := []string{
        `CREATE TABLE users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);`,
        `CREATE TABLE IF NOT EXISTS todos (
            id INT AUTO_INCREMENT PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            description TEXT,
            completed BOOLEAN DEFAULT FALSE,
            user_id INT,
            FOREIGN KEY (user_id) REFERENCES users(id)
        );`,
    }

    for _, query := range queries {
        _, err := db.Exec(query)
        if err != nil {
            log.Fatalf("Error running migration query: %v", err)
        }
    }

    log.Println("Migrations executed successfully!")
}