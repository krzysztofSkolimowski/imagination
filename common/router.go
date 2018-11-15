package common

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	mux *mux.Router
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func NewRouter() *Router {
	r := mux.NewRouter()
	return &Router{r}
}

func (r *Router) Route(path string, f func(w http.ResponseWriter, r *http.Request), methods ...string) {
	r.mux.HandleFunc(path, f).Methods(methods...)
}
