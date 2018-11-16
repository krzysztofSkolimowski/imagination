package image

import (
	"fmt"
	"io"
)

type Transform string

type fileService interface {
	SaveFile(fileID string, r io.Reader) error
	LoadFile(fileID string) (io.ReadCloser, int, error)
	DeleteFile(fileID string) error
	DeleteDirectory(directory string) error
	CopyFile(srcFileID string, destFileID string, public bool) error
}

//type accessControlFileService interface {
//	fileService
//	SavePrivateFile(fileID string, r io.Reader) error
//}

type resolver interface {
	Resolve(fileID string) string
}

type Service struct {
	pathResolver, urlResolver      resolver
	localFileService, cloudStorage fileService
	availableTransforms            map[Transform]interface{}
}

func NewService(
	pathResolver, urlResolver resolver,
	localFileService, cloudStorage fileService,
	availableTransforms map[Transform]interface{},
) *Service {
	return &Service{
		pathResolver, urlResolver,
		localFileService, cloudStorage,
		availableTransforms,
	}
}

type ProcessCmd struct {
	File       []byte
	FileName   string
	transforms []Transform
}

func (s Service) Process(cmd ProcessCmd) error {
	//todo - validate file
	fmt.Println("=============================")
	fmt.Println("processing image")
	fmt.Println("=============================")
	//todo - check transforms
	//todo - saveFile on temporary path
	//todo - defer remove file
	//todo - decode image
	return nil
}
