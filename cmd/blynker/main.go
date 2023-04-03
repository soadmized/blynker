package main

import (
	"blynker/api"
	"blynker/internal/config"
	"log"
	"net/http"
	"strconv"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	addr := ":" + strconv.Itoa(conf.AppPort)
	server := api.New(conf)
	err = http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatal(err)
	}
}
