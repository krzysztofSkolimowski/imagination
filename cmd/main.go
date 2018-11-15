package main

import (
	"context"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	config := Config{}
	config.LoadConfig()

	//todo - give proper abstraction to logger
	logger := log.New()

	r := mux.NewRouter()

	logger.Println("Service started")

	logger.Fatal(http.ListenAndServe(":8080", r))
}

func runServer(config Config) {
	ctx := context.Background()
	//m, cancel := 
}
