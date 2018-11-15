package rest

import (
	"fmt"
	"github.com/krzysztofSkolimowski/imagination/cmd/modules"
	"github.com/krzysztofSkolimowski/imagination/common"
	"net/http"
)

func AddImageResource(svc *modules.ImaginationServices, r *common.Router) {
	r.Route("/", releaseYourImagination(svc), "GET")
}

//debug method written in order to check if server works
func releaseYourImagination(svc *modules.ImaginationServices) func(w http.ResponseWriter, r *http.Request) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "release your imagination")
	}
	return handleFunc
}
