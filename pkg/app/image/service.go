package image

import (
	"fmt"
	"io"
)

type Transform string

type LocalFileService interface {
	SaveFile(fileID string, r io.Reader) error
	LoadFile(fileID string) (io.ReadCloser, int, error)
	DeleteFile(fileID string) error
	DeleteDirectory(directory string) error
	CopyFile(srcFileID string, destFileID string, public bool) error
}

type CloudStorage interface {
	SaveFile(fileID string, r io.Reader) error
	LoadFile(fileID string) (io.ReadCloser, int, error)
	DeleteFile(fileID string) error
	DeleteDirectory(directory string) error
	CopyFile(srcFileID string, destFileID string, public bool) error
}

type PathResolver interface {
	Resolve(fileID string) string
}

type URLResolver interface {
	Resolve(fileID string) string
}

type Service struct {
	pathResolver        PathResolver
	urlResolver         URLResolver
	localFileService    LocalFileService
	cloudStorage        CloudStorage
	availableTransforms map[Transform]struct{}
}

func NewService(
	pathResolver PathResolver,
	urlResolver URLResolver,
	localFileService LocalFileService,
	cloudStorage CloudStorage,
	available []Transform,
) *Service {
	transforms := make(map[Transform]struct{}, len(available))

	for _, v := range available {
		transforms[v] = struct{}{}
	}

	return &Service{
		pathResolver,
		urlResolver,
		localFileService,
		cloudStorage,
		transforms,
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
