package image

import (
	"io"
)

type transform string

type fileService interface {
	SaveFile(fileID string, r io.Reader) error
	LoadFile(fileID string) (io.ReadCloser, int, error)
	DeleteFile(fileID string) error
	DeleteDirectory(directory string) error
	CopyFile(srcFileID string, destFileID string, public bool) error
}

type accessControlFileService interface {
	fileService
	SavePrivateFile(fileID string, r io.Reader) error
}

type resolver interface {
	Resolve(fileID string) string
}

type Service struct {
	resolver            resolver
	fileService         fileService
	availableTransforms map[transform]interface{}
}

func NewService(r resolver, fs fileService, transforms map[transform]interface{}) Service {
	return Service{r, fs, transforms}
}

type ProcessCmd struct {
	//todo - refactor to uuids
	ID         int64
	File       []byte
	transforms []transform
}

func (s Service) Process(cmd ProcessCmd) error {
	//todo - validate file
	//todo - check transforms
	//todo - saveFile on temporary path
	//todo - defer remove file
	//todo - decode image
	return nil
}
