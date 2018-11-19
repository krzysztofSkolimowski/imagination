package image

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"path"
)

const maxFileNameLength = 255

type Transform string
type Format string

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
	availableFormats    map[Format]struct{}
}

func NewService(
	pathResolver PathResolver,
	urlResolver URLResolver,
	localFileService LocalFileService,
	cloudStorage CloudStorage,
	available []Transform,
	availableFormats []Format,
) *Service {
	transforms := make(map[Transform]struct{}, len(available))
	for _, v := range available {
		transforms[v] = struct{}{}
	}

	formats := make(map[Format]struct{}, len(availableFormats))
	for _, v := range availableFormats {
		formats[v] = struct{}{}
	}

	return &Service{
		pathResolver,
		urlResolver,
		localFileService,
		cloudStorage,
		transforms,
		formats,
	}
}

type ProcessCmd struct {
	File        []byte
	FileName    string
	SaveToCloud bool
	Transforms  []Transform
}

func (s Service) Process(cmd ProcessCmd) error {
	fileName := sanitize(cmd.FileName)
	temporaryPath := s.pathResolver.Resolve(fileName)
	if err := s.localFileService.SaveFile(temporaryPath, bytes.NewBuffer(cmd.File)); err != nil {
		return err
	}

	_, sourceFormat, err := image.Decode(bytes.NewBuffer(cmd.File))
	if err != nil {
		return errors.Wrap(err, "Cannot decode provided image")
	}

	if _, ok := s.availableFormats[Format(sourceFormat)]; !ok {
		return errors.New("Unsupported format")
	}

	//todo - check transforms
	for _, v := range cmd.Transforms {
		if _, ok := s.availableTransforms[v]; ok {
			fmt.Println("performing transform: ", v)
		}
	}
	//todo - defer remove file
	//todo - decode image

	//todo - save to cloud or stream back
	//paths := s.urlResolver.Resolve(fileName)
	if cmd.SaveToCloud {
		return s.cloudStorage.SaveFile(fileName, bytes.NewBuffer(cmd.File))
	}
	return nil
}

func sanitize(fileName string) string {
	fileName = path.Base(path.Clean(fileName))
	maxLen := maxFileNameLength
	if len(fileName) > maxLen {
		fileName = fileName[len(fileName)-maxLen:]
	}
	return fileName
}
