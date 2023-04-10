package main

import (
	"log"
	"strconv"

	"blynker/api"
	"blynker/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	addr := ":" + strconv.Itoa(conf.AppPort)
	server := api.New(conf)

	err = server.Router.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}
