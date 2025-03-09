package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/matheusteixeira7/ecommerce-challenge/packages/user-service/config"
	"github.com/matheusteixeira7/ecommerce-challenge/packages/user-service/infra/database"

	"log"
	"net/http"
)

func main() {
	env, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// open db connection
	conn, err := database.OpenConn(env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_PASSWORD, env.DB_NAME, env.DB_SSL_MODE)
	if err != nil {
		panic(err)
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing the database connection: %v", err)
		}
	}(conn)

	// run migrations
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create driver: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://infra/database/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Successfully run migrations.")

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Printf("Starting server on port %s", env.PORT)
	err = http.ListenAndServe(":"+env.PORT, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
