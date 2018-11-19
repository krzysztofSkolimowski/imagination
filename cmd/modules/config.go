package modules

import (
	"github.com/krzysztofSkolimowski/imagination/pkg/infrastructure/files"
	"os"
)

type Config struct {
	port              string
	uploadsDir        string
	logsDir           string
	s3Bucket          string
	s3Region          string
	s3BaseURL         string
	s3AccessKeyId     string
	s3SecretAccessKey string
	s3MinioEnabled    bool
	s3MinioUrl        string
	mysqlPort         string
	mysqlUser         string
	mysqlPassword     string
}

func (c *Config) LoadConfig() {
	c.port = os.Getenv("IMAGINATION_PORT")
	c.uploadsDir = os.Getenv("IMAGINATION_UPLOADS_DIR")
	c.logsDir = os.Getenv("IMAGINATION_LOGS_DIR")
	c.s3Bucket = os.Getenv("IMAGINATION_AWS_S3_BUCKET")
	c.s3Region = os.Getenv("IMAGINATION_AWS_S3_REGION")
	c.s3BaseURL = os.Getenv("IMAGINATION_AWS_S3_BASE_URL")
	c.s3AccessKeyId = os.Getenv("IMAGINATION_AWS_ACCESS_KEY_ID")
	c.s3SecretAccessKey = os.Getenv("IMAGINATION_AWS_SECRET_ACCESS_KEY")
	if os.Getenv("IMAGINATION_AWS_S3_MINIO_ENABLED") == "1" {
		c.s3MinioEnabled = true
	}
	c.s3MinioUrl = os.Getenv("IMAGINATION_AWS_S3_MINIO_URL")
	c.mysqlPort = os.Getenv("IMAGINATION_MYSQL_PORT")
	c.mysqlUser = os.Getenv("IMAGINATION_MYSQL_USER")
	c.mysqlPassword = os.Getenv("IMAGINATION_MYSQL_PASSWORD")
}

func (c Config) Port() string                                { return c.port }
func (c Config) UploadsDir() files.UploadsDir                { return files.UploadsDir(c.uploadsDir) }
func (c Config) S3BaseURL() files.BaseURL                    { return files.BaseURL(c.s3BaseURL) }
func (c Config) S3Bucket() files.Bucket                      { return files.Bucket(c.s3Bucket) }
func (c Config) S3Region() files.AWSRegion                   { return files.AWSRegion(c.s3Region) }
func (c Config) S3AccessKeyID() files.AWSAccessKeyID         { return files.AWSAccessKeyID(c.s3AccessKeyId) }
func (c Config) S3SecretAccessKey() files.AWSSecretAccessKey { return files.AWSSecretAccessKey(c.s3SecretAccessKey) }
func (c Config) S3MinioEnabled() files.MinioEnabled          { return files.MinioEnabled(c.s3MinioEnabled) }
func (c Config) S3MinioUrl() files.MinioURL                  { return files.MinioURL(c.s3MinioUrl) }
