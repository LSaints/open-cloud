package main

import (
	config "backend/configs"
	"backend/pkg/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	router := router.Initialize()

	fmt.Printf("Listening on: %s:%d", config.Host, config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
