package modules

import "os"

type Config struct {
	Port       string
	UploadsDir string
	LogsDir    string

	S3Bucket          string
	S3Region          string
	S3Base_url        string
	S3AccessKeyId     string
	S3SecretAccessKey string

	S3MinioEnabled string
	S3MinioUrl     string

	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
}

func (c *Config) LoadConfig() {
	c.Port = os.Getenv("IMAGINATION_PORT")
	c.UploadsDir = os.Getenv("IMAGINATION_UPLOADS_DIR")
	c.LogsDir = os.Getenv("IMAGINATION_LOGS_DIR")
	c.S3Bucket = os.Getenv("IMAGINATION_AWS_S3_BUCKET")
	c.S3Region = os.Getenv("IMAGINATION_AWS_S3_REGION")
	c.S3Base_url = os.Getenv("IMAGINATION_AWS_S3_BASE_URL")
	c.S3AccessKeyId = os.Getenv("IMAGINATION_AWS_ACCESS_KEY_ID")
	c.S3SecretAccessKey = os.Getenv("IMAGINATION_AWS_SECRET_ACCESS_KEY")
	c.S3MinioEnabled = os.Getenv("IMAGINATION_AWS_S3_MINIO_ENABLED")
	c.S3MinioUrl = os.Getenv("IMAGINATION_AWS_S3_MINIO_URL")
	c.MysqlPort = os.Getenv("IMAGINATION_MYSQL_PORT")
	c.MysqlUser = os.Getenv("IMAGINATION_MYSQL_USER")
	c.MysqlPassword = os.Getenv("IMAGINATION_MYSQL_PASSWORD")
}
