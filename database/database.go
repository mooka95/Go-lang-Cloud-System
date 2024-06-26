package database

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/lib/pq"
)

var (
    DB *sql.DB
    connectionString = "postgres://postgres:Torm22torm*@localhost:5432/cloudsystem?sslmode=disable"
)

// Init initializes the database connection pool
func Init() {
    var err error
    DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Panicf("Unable to connect to database: %v\n", err)
    }

    // Configure the connection pool
    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)
    DB.SetConnMaxLifetime(30 * time.Minute)

    // Verify the connection
    if err = DB.Ping(); err != nil {
        log.Fatalf("Unable to verify connection: %v\n", err)
    }

    fmt.Println("Database connection pool established")
}
