package main

import (
	"context"
	"github.com/krzysztofSkolimowski/imagination/cmd/modules"
	"github.com/krzysztofSkolimowski/imagination/common"
	"github.com/krzysztofSkolimowski/imagination/common/middleware"
	"github.com/krzysztofSkolimowski/imagination/pkg/app/image"
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
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	services, err := modules.SetupServices(
		config.UploadsDir(),
		config.S3BaseURL(),
		nil,
		config.S3Region(),
		config.S3AccessKeyID(),
		config.S3SecretAccessKey(),
		config.S3MinioEnabled(),
		config.S3MinioUrl(),
		config.S3Bucket(),
		[]image.Transform{
			"crop",
			"rot_90",
			"rot_180",
			"rm_exif",
		},
	)
	if err != nil {
		panic(err)
	}

	r := common.NewRouter()
	rest.AddImageResource(services, r)

	log.Info("Starting imagination")
	return http.ListenAndServe(
		":"+config.Port(),
		middleware.RequestLogger(*services.Logger)(r),
	)

}
