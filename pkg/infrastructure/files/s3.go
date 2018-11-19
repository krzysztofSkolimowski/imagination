package files

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"net/url"
	"path"
	"strings"

	"gopkg.in/h2non/filetype.v1"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

var ErrInvalidDirectory = errors.New("invalid directory")

type Bucket string

type S3FileService struct {
	uploader *s3manager.Uploader
	service  *s3.S3
	bucket   Bucket
	//todo - replace logger with proper abstraction
	logger logrus.Logger
}

func NewS3FileService(u *s3manager.Uploader, service *s3.S3, bucket Bucket, l *logrus.Logger) S3FileService {
	return S3FileService{u, service, bucket, *l}
}

func (fs S3FileService) SaveFile(fileID string, r io.Reader) error {
	return fs.saveFile(fileID, r, true)
}

func (fs S3FileService) saveFile(fileID string, r io.Reader, public bool) error {
	var headerBuf bytes.Buffer
	tee := io.TeeReader(r, &headerBuf)

	mimeType, err := filetype.MatchReader(tee)
	if err != nil {
		return errors.Wrap(err, "problem with parsing mime type")
	}

	contentType := mimeType.MIME.Value

	fs.logger.WithFields(logrus.Fields{
		"file":         fileID,
		"content_type": contentType,
		"public":       public,
	}).Info("Uploading file to S3")

	fileContent := io.MultiReader(&headerBuf, r)
	uploadInput := &s3manager.UploadInput{
		Bucket:      aws.String(string(fs.bucket)),
		Key:         aws.String(fileID),
		Body:        fileContent,
		ContentType: aws.String(contentType),
	}

	if public {
		uploadInput.ACL = aws.String("public-read")
	}

	_, err = fs.uploader.Upload(uploadInput)
	if err != nil {
		return errors.Wrap(err, "failed to upload file")
	}

	return err
}

func (fs S3FileService) LoadFile(fileID string) (io.ReadCloser, int, error) {
	params := &s3.GetObjectInput{
		Bucket: aws.String(string(fs.bucket)),
		Key:    aws.String(fileID),
	}

	fs.logger.WithFields(logrus.Fields{
		"file": fileID,
	}).Info("Downloading file from S3")

	object, err := fs.service.GetObject(params)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to download file")
	}

	contentLength := 0
	if object.ContentLength != nil {
		contentLength = int(*object.ContentLength)
	}

	return object.Body, contentLength, nil
}

func (fs S3FileService) DeleteFile(fileID string) error {
	params := &s3.DeleteObjectInput{
		Bucket: aws.String(string(fs.bucket)),
		Key:    aws.String(fileID),
	}

	fs.logger.WithFields(logrus.Fields{
		"file": fileID,
	}).Info("Deleting file from S3")

	_, err := fs.service.DeleteObject(params)

	if err != nil {
		return errors.Wrap(err, "failed to delete file")
	}

	return err
}

// DeleteDirectory deletes all objects which keys begin with the specified directory.
// directory must end with an "/".
func (fs S3FileService) DeleteDirectory(directory string) error {
	directory = strings.TrimSpace(directory)
	err := validateDirectory(directory)
	if err != nil {
		return err
	}

	files, err := fs.findFilesWithPrefix(directory)
	if err != nil {
		return err
	}

	var objectsToRemove []string
	for _, file := range files {
		// Sanity check
		if !strings.HasPrefix(*file.Key, directory) {
			return errors.Errorf("invalid prefix for file %s in directory %s", *file.Key, directory)
		}

		objectsToRemove = append(objectsToRemove, *file.Key)
	}

	// S3 supports batch delete for up to 1000 keys
	return processChunks(objectsToRemove, 1000, fs.deleteFiles)
}

func validateDirectory(dir string) error {
	if dir == "" || dir == "/" || !strings.HasSuffix(dir, "/") {
		return ErrInvalidDirectory
	}
	return nil
}

func (fs S3FileService) findFilesWithPrefix(prefix string) ([]*s3.Object, error) {
	params := &s3.ListObjectsInput{
		Bucket: aws.String(string(fs.bucket)),
		Prefix: aws.String(prefix),
	}

	response, err := fs.service.ListObjects(params)
	if err != nil {
		return nil, errors.Wrap(err, "failed listing files")
	}

	return response.Contents, nil
}

func (fs S3FileService) deleteFiles(files []string) error {
	var objects []*s3.ObjectIdentifier
	for _, file := range files {
		objects = append(objects, &s3.ObjectIdentifier{
			Key: aws.String(file),
		})
		fs.logger.WithFields(logrus.Fields{
			"file": file,
		}).Debug("Deleting file from S3")
	}

	deleteParams := &s3.DeleteObjectsInput{
		Bucket: aws.String(string(fs.bucket)),
		Delete: &s3.Delete{
			Objects: objects,
		},
	}

	_, err := fs.service.DeleteObjects(deleteParams)
	if err != nil {
		return errors.Wrap(err, "failed deleting files")
	}

	return nil
}

func processChunks(objects []string, maxChunkSize int, processor func([]string) error) error {
	for len(objects) > 0 {
		chunkSize := len(objects)
		if chunkSize > maxChunkSize {
			chunkSize = maxChunkSize
		}

		chunk := objects[:chunkSize]
		objects = objects[chunkSize:]

		err := processor(chunk)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fs S3FileService) CopyFile(srcFileID string, destinationFileID string, public bool) error {
	params := &s3.CopyObjectInput{
		Bucket:     aws.String(string(fs.bucket)),
		Key:        aws.String(destinationFileID),
		CopySource: aws.String(url.QueryEscape(path.Join(string(fs.bucket), srcFileID))),
	}

	if public {
		params.ACL = aws.String("public-read")
	}

	fs.logger.WithFields(logrus.Fields{
		"source_file": srcFileID,
		"dest_file":   destinationFileID,
		"public":      public,
	}).Debug("Copying file on S3")

	if _, err := fs.service.CopyObject(params); err != nil {
		return errors.Wrap(err, "failed to copy file")
	}

	return nil
}
