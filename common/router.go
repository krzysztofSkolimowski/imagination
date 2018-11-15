package common

import (
	"github.com/gorilla/mux"
	"net/http"
)

//todo - create proper abstraction in common
type Router struct {
	mux *mux.Router
}

func (r *Router) Route(path string, f func(w http.ResponseWriter, r *http.Request), methods ...string) {
	r.mux.HandleFunc(path, f).Methods(methods...)
}
