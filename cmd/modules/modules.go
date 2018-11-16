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
	Common commonServices
	Infra  infraServices
	App    appServices
}

type commonServices struct {
	Logger       *logrus.Logger
	URLResolver  *files.URLResolver
	PathResolver *files.PathResolver
}

type infraServices struct {
	S3Service  *s3.S3
	S3Uploader *s3manager.Uploader
}

type appServices struct {
	ImageService *image.Service
}

func SetupServices(ctx context.Context, config Config) (*Services, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	services := &Services{}

	services.initializeCommon(ctx, config)
	services.initializeInfrastructure(ctx, config)
	services.initializeApp(ctx, config)

	return services, cancel
}

func (svc *Services) initializeCommon(ctx context.Context, config Config) {
	svc.Common.Logger = logrus.New()

	pathResolver, err := files.NewPathResolver(config.UploadsDir)
	if err != nil {
		panic(err)
	}
	svc.Common.PathResolver = &pathResolver

	urlResolver, err := files.NewURLResolver(config.S3Base_url)
	if err != nil {
		panic(err)
	}
	svc.Common.URLResolver = &urlResolver
}

//todo check nil pointer
//func (svc *Services) initializeS3(config Config) {
//
//
//	fmt.Println("...................")
//	fmt.Println(uploader)
//	fmt.Println(service)
//	svc.S3.Uploader = uploader
	//svc.S3.Uploader = uploader
	//svc.S3.Service = service
	//svc.S3 = struct {
	//}{}
	//fmt.Println(svc.S3.Uploader)
	//fmt.Println("...................")
//}

func (svc *Services) initializeInfrastructure(ctx context.Context, config Config) {
	awsConfig := &aws.Config{
		Region: aws.String(config.S3Region),
		Credentials: credentials.NewStaticCredentials(
			config.S3AccessKeyId,
			config.S3SecretAccessKey,
			"",
		),
	}

	if config.S3MinioEnabled {
		awsConfig.Endpoint = aws.String(config.S3MinioUrl)
		awsConfig.DisableSSL = aws.Bool(true)
		awsConfig.S3ForcePathStyle = aws.Bool(true)
	}

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		panic(err)
	}
//	todo - initilize s3

}

func (svc *Services) initializeApp(ctx context.Context, config Config) {
	svc.App.ImageService = image.NewService(
		svc.Common.PathResolver, svc.Common.URLResolver,
		files.NewLocalFileService(config.UploadsDir, []files.Validator{}),
		files.NewS3FileService(svc.Infra.S3Uploader, svc.Infra.S3Service, config.S3Bucket, *svc.Common.Logger),
		map[image.Transform]interface{}{
			image.Transform("crop"):    nil,
			image.Transform("rot_90"):  nil,
			image.Transform("rot_180"): nil,
			image.Transform("rm_exif"): nil,
		},
	)

}
