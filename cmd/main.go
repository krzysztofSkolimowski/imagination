package main

import (
	"context"
	"github.com/krzysztofSkolimowski/imagination/cmd/modules"
	"github.com/krzysztofSkolimowski/imagination/common"
	"github.com/krzysztofSkolimowski/imagination/common/middleware"
	"github.com/krzysztofSkolimowski/imagination/pkg/interfaces/rest"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	config := modules.Config{}
	config.LoadConfig()
	log.Fatal(runServer(config))
}

func runServer(config modules.Config) error {
	ctx := context.Background()

	services, cancel := modules.SetupServices(ctx, config)
	defer cancel()

	r := common.NewRouter()
	rest.AddImageResource(services, r)

	log.Info("Starting imagination")
	return 	http.ListenAndServe(
		":" + config.Port,
		middleware.RequestLogger(*services.Logger)(r),
	)

}
