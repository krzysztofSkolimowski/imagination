package modules

import (
	"context"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/krzysztofSkolimowski/imagination/pkg/app/image"
	"github.com/sirupsen/logrus"
)

type ImaginationServices struct {
	Logger *logrus.Logger

	S3 *struct {
		Service  *s3.S3
		Uploader *s3manager.Uploader
	}

	ImageService *image.Service
}

func SetupImaginationServices(ctx context.Context) (*ImaginationServices, context.CancelFunc){
	ctx, cancel := context.WithCancel(ctx)
	services := ImaginationServices{}

	logger := logrus.New()

	services.Logger = logger

	return &services, cancel
}
