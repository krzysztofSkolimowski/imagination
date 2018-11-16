package modules

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/krzysztofSkolimowski/imagination/pkg/app/image"
	"github.com/krzysztofSkolimowski/imagination/pkg/infrastructure/files"
	"github.com/sirupsen/logrus"
)

type Services struct {
	Logger *logrus.Logger

	S3 *struct {
		Service  *s3.S3
		Uploader *s3manager.Uploader
	}

	//URLResolver files.URLResolver

	ImageService *image.Service
}

func SetupImaginationServices(ctx context.Context, config Config) (*Services, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	services := &Services{}

	logger := logrus.New()

	services.Logger = logger

	pathResolver, err := files.NewPathResolver(config.UploadsDir)
	if err != nil {
		panic(err)
	}

	urlResolver, err := files.NewURLResolver(config.S3Base_url)
	if err != nil {
		panic(err)
	}

	services.initializeS3(config)

	services.ImageService = image.NewService(
		pathResolver, urlResolver,
		files.NewLocalFileService(config.UploadsDir, []files.Validator{}),
		files.NewS3FileService(services.S3.Uploader, services.S3.Service, config.S3Bucket, *services.Logger),
		map[image.Transform]interface{}{
			image.Transform("crop"):    nil,
			image.Transform("rot_90"):  nil,
			image.Transform("rot_180"): nil,
			image.Transform("rm_exif"): nil,
		},
	)

	return services, cancel
}

//todo check nil pointer
func (svc *Services) initializeS3(config Config) {
	c := &aws.Config{
		Region: aws.String(config.S3Region),
		Credentials: credentials.NewStaticCredentials(
			config.S3AccessKeyId,
			config.S3SecretAccessKey,
			"",
		),
	}

	if config.S3MinioEnabled {
		c.Endpoint = aws.String(config.S3MinioUrl)
		c.DisableSSL = aws.Bool(true)
		c.S3ForcePathStyle = aws.Bool(true)
	}

	sess, err := session.NewSession(c)
	if err != nil {
		panic(err)
	}

	svc.S3.Uploader = s3manager.NewUploader(sess)
	svc.S3.Service = s3.New(sess)
}
