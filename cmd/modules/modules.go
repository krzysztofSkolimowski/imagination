package modules

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
	"imagination/common"
	"imagination/pkg/app/image"
)

type ImaginationServices struct {
	Logger *logrus.Logger
	Router *common.Router

	S3     *struct {
		Service  *s3.S3
		Uploader *s3manager.Uploader
	}

	ImageService *image.Service
}
