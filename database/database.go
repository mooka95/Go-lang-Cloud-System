package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

// Init initializes the database connection pool and creates tables if they do not exist
func Init() {
	// Get environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Build the connection string
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

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

	// Create tables if they do not exist
	if err := createTables(); err != nil {
		log.Fatalf("Error creating tables: %v\n", err)
	}

	fmt.Println("Tables created successfully")
}

// createTables creates necessary tables if they do not exist
func createTables() error {
	createTableStmt := `
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    CREATE TYPE IF NOT EXISTS os_type AS ENUM ('Linux', 'Windows', 'MacOS');
        CREATE TABLE IF NOT EXISTS public.users (
            id SERIAL PRIMARY KEY,
            first_name VARCHAR(20) NOT NULL,
            last_name VARCHAR(20) NOT NULL,
            email VARCHAR(255) NOT NULL,
            password VARCHAR(255) NOT NULL,
            identifier VARCHAR(255) DEFAULT uuid_generate_v4() NOT NULL,
            CONSTRAINT users_email_key UNIQUE (email),
            CONSTRAINT users_identifier_key UNIQUE (identifier)
        );

        CREATE TABLE IF NOT EXISTS public.addresses (
            id SERIAL PRIMARY KEY,
            city VARCHAR(255) NOT NULL,
            street VARCHAR(255) NOT NULL,
            country VARCHAR(255) NOT NULL,
            identifier VARCHAR(255) NOT NULL,
            user_id INT NOT NULL,
            CONSTRAINT unique_address_user_id UNIQUE (city, street, country, user_id),
            CONSTRAINT address_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id)
        );

        CREATE TABLE IF NOT EXISTS public.firewalls (
            id SERIAL PRIMARY KEY,
            name VARCHAR(20) NOT NULL,
            identifier VARCHAR(255) NOT NULL,
            user_id INT NULL,
            CONSTRAINT firewalls_identifier_key UNIQUE (identifier),
            CONSTRAINT unique_firewall_name_user_id UNIQUE (name, user_id),
            CONSTRAINT firewalls_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id)
        );

        CREATE TABLE IF NOT EXISTS public.virtualmachines (
            id SERIAL PRIMARY KEY,
            hostname VARCHAR(255) NOT NULL,
            operating_system public.os_type NULL,
            is_active BOOLEAN NOT NULL,
            identifier VARCHAR(255) DEFAULT uuid_generate_v4() NOT NULL,
            user_id INT NOT NULL,
            CONSTRAINT virtualmachines_identifier_key UNIQUE (identifier),
            CONSTRAINT virtualmachines_unique UNIQUE (hostname),
            CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id)
        );

        CREATE TABLE IF NOT EXISTS public.virtualmachines_firewalls (
            id SERIAL PRIMARY KEY,
            identifier VARCHAR(255) NOT NULL,
            virtualmachine_id INT NOT NULL,
            firewall_id INT NOT NULL,
            CONSTRAINT unique_combination UNIQUE (virtualmachine_id, firewall_id),
            CONSTRAINT firewall_id_fkey FOREIGN KEY (firewall_id) REFERENCES public.firewalls(id) ON DELETE CASCADE,
            CONSTRAINT virtualmachine_id_fkey FOREIGN KEY (virtualmachine_id) REFERENCES public.virtualmachines(id) ON DELETE CASCADE
        );
    `

	_, err := DB.Exec(createTableStmt)
	if err != nil {
		return err
	}

	return nil
}
