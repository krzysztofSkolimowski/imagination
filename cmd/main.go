package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Release your imagination")
	}).Methods("GET")
	fmt.Println("Service started")

	log.Fatal(http.ListenAndServe(":8080", r))
}
