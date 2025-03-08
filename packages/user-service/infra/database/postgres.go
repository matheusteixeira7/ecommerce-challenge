package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func OpenConn(host, port, user, password, dbname, sslmode string) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	var db *sql.DB
	var err error

	maxAttempts := 5
	waitInterval := 5 * time.Second

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("Failed to open database connection: %v", err)
			time.Sleep(waitInterval)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Attempt %d: Database is not ready yet, retrying in %v...", attempt, waitInterval)
			time.Sleep(waitInterval)
		} else {
			log.Println("Successfully connected to the database.")
			return db, nil
		}
	}

	return nil, fmt.Errorf("after %d attempts, failed to connect to database: %v", maxAttempts, err)
}
