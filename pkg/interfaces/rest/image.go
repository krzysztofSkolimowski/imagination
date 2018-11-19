package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/krzysztofSkolimowski/imagination/cmd/modules"
	"github.com/krzysztofSkolimowski/imagination/common"
	"github.com/krzysztofSkolimowski/imagination/pkg/app/image"
	"io"
	"net/http"
)

func AddImageResource(svc *modules.Services, r *common.Router) {
	r.Route("/", releaseYourImagination(svc), http.MethodGet)
	r.Route("/image", processImage(svc), http.MethodPost)
}

//debug method written in order to check if server works
func releaseYourImagination(svc *modules.Services) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "release your imagination")
	}
}

func processImage(svc *modules.Services) func(w http.ResponseWriter, r *http.Request) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var request PostProcessImageRequest
		err := decoder.Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd := image.ProcessCmd{
			Transforms:  request.Transforms,
			SaveToCloud: request.SaveToCloud,
			ImageURL:    request.ImageURL,
		}


		f, err := svc.ImageService.Process(cmd)
		if  err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		io.Copy(w, bytes.NewBuffer(f))
	}
	return handleFunc
}
