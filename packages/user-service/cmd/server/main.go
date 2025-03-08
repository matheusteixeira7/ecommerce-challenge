package main

import (
	"database/sql"
	"fmt"
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
	conn, err := database.OpenConn(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode)
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
		"file://../../infra/database/migrations",
		"postgres", driver)
	m.Up()
	if err != nil {
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
