package main

import (
	"fmt"
	"github.com/matheusteixeira7/ecommerce-challenge/packages/user-service/config"
	"log"
	"net/http"
)

func main() {
	env, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

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
