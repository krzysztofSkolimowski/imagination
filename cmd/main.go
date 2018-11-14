package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	//todo - give proper abstraction to logger
	logger := log.New()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Release your imagination")
	}).Methods("GET")
	logger.Println("Service started")

	logger.Fatal(http.ListenAndServe(":8080", r))
}
