// +build

package modules

import (
	"github.com/google/go-cloud/wire"
	"github.com/krzysztofSkolimowski/imagination/pkg/app/image"
	"github.com/krzysztofSkolimowski/imagination/pkg/infrastructure/files"
	"github.com/sirupsen/logrus"
)


var providers = wire.NewSet(
	NewServices,

	logrus.New,

	image.NewService,

	files.NewPathResolver,
	wire.Bind(new(image.PathResolver), files.PathResolver{}),

	files.NewURLResolver,
	wire.Bind(new(image.URLResolver), files.URLResolver{}),

	files.NewLocalFileService,
	wire.Bind(new(image.LocalFileService), files.LocalFileService{}),

	files.NewS3FileService,
	wire.Bind(new(image.CloudStorage), files.S3FileService{}),

	files.NewS3Config,
	files.NewAWSSession,
	files.NewS3Uploader,
	files.NewS3Service,
)

func SetupServices(
	uploadsDir files.UploadsDir,
	baseURL files.BaseURL,
	validators []files.Validator,
	region files.AWSRegion,
	accessKeyID files.AWSAccessKeyID,
	secretAccessKey files.AWSSecretAccessKey,
	minioEnabled files.MinioEnabled,
	minioURL files.MinioURL,
	bucket files.Bucket,
	transforms []image.Transform,
) (*Services, error) {
	wire.Build(providers)
	return nil, nil
}
