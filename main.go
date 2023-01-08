package main

import (
	"log"
	"net/http"
	"os"

	"blynker/internal/api"
)

func main() {
	port := os.Getenv("PORT")
	addr := ":" + port
	server := api.New()

	err := http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatal(err)
	}
}
