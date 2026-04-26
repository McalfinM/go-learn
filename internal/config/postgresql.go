package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connStr)

	log.Printf("Connecting with: %s\n", connStr)

	if connStr == "" {
		log.Fatal("DATABASE_URL is empty. Make sure it's set in your environment.")
	}

	if err != nil {
		log.Fatal("failed to connect db:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("db not reachable:", err)
	}

	log.Println("Database connected")
	return db
}
