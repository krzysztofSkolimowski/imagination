package rest

import (
	"fmt"
	"imagination/common"
	"net/http"
)

func AddImageResource(r *common.Router) {
	//r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//
	//	fmt.Fprint(w, "Release your imagination")
	//}).Methods("GET")

	r.Route("/", releaseYourImagination, "GET")

}

func releaseYourImagination(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Release your imagination")
}
