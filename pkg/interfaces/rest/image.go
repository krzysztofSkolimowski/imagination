package rest

import (
	"fmt"
	"github.com/krzysztofSkolimowski/imagination/cmd/modules"
	"github.com/krzysztofSkolimowski/imagination/common"
	"github.com/krzysztofSkolimowski/imagination/pkg/app/image"
	"io/ioutil"
	"net/http"
)

func AddImageResource(svc *modules.Services, r *common.Router) {
	r.Route("/", releaseYourImagination(svc), http.MethodGet)
	r.Route("/image", uploadImage(svc), http.MethodPost)
}

//debug method written in order to check if server works
func releaseYourImagination(svc *modules.Services) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "release your imagination")
	}
}

func uploadImage(svc *modules.Services) func(w http.ResponseWriter, r *http.Request) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		formFile, fileHeader, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer func() {
			if err := formFile.Close(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()

		f, err := ioutil.ReadAll(formFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cmd := image.ProcessCmd{
			File:     f,
			FileName: fileHeader.Filename,
		}

		if err := svc.ImageService.Process(cmd); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	return handleFunc
}

func transform(svc *modules.Services) func(w http.ResponseWriter, r *http.Request) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {

	}
	return handleFunc
}
