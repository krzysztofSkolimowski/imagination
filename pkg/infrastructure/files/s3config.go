package files

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AWSRegion string
type AWSAccessKeyID string
type AWSSecretAccessKey string
type MinioEnabled bool
type MinioURL string

type s3Config struct {
	Region          AWSRegion
	AccessKeyID     AWSAccessKeyID
	SecretAccessKey AWSSecretAccessKey
	MinioEnabled    MinioEnabled
	MinioURL        MinioURL
}

func NewS3Config(
	region AWSRegion,
	key AWSAccessKeyID,
	secret AWSSecretAccessKey,
	minioEnabled MinioEnabled,
	minioURL MinioURL,
) s3Config {
	return s3Config{
		region,
		key,
		secret,
		minioEnabled,
		minioURL,
	}

}

func NewAWSSession(c s3Config) (*session.Session, error) {
	awsConfig := &aws.Config{
		Region: aws.String(string(c.Region)),
		Credentials: credentials.NewStaticCredentials(
			string(c.AccessKeyID),
			string(c.SecretAccessKey),
			"",
		),
	}

	if bool(c.MinioEnabled) {
		awsConfig.Endpoint = aws.String(string(c.MinioURL))
		awsConfig.DisableSSL = aws.Bool(true)
		awsConfig.S3ForcePathStyle = aws.Bool(true)
	}

	return session.NewSession(awsConfig)
}

func NewS3Uploader(s *session.Session) *s3manager.Uploader {
	return s3manager.NewUploader(s)
}

func NewS3Service(s *session.Session) *s3.S3 {
	return s3.New(s)
}
