package main

import (
	"handlers"
	"log"
	"net/http"
)

func main() {
	handlers.HandleAllHttpFuncs()

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
	log.Fatal(err)
}
