package utils

import (
    "fmt"
    "log"
    "os"
    "sync"

    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
	
)

var (
    db   *sqlx.DB
    once sync.Once
)

func GetDB() *sqlx.DB {
    once.Do(func() {
        dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
            os.Getenv("DB_USER"),
            os.Getenv("DB_PASSWORD"),
            os.Getenv("DB_HOST"),
            os.Getenv("DB_PORT"),
            os.Getenv("DB_NAME"),
        )

        var err error
        db, err = sqlx.Connect("mysql", dsn)
        if err != nil {
            log.Fatalf("Error connecting to the database: %v", err)
        }
    })

    return db
}
