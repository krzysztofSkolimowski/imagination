package image

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/http"
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
	ImageURL    string
	SaveToCloud bool
	Transforms  []string
}

func (s Service) Process(cmd ProcessCmd) ([]byte, error) {
	imageURL := cmd.ImageURL
	fileName := sanitize(imageURL)

	if err := s.DownloadFile(fileName, imageURL); err != nil {
		return nil, errors.Wrap(err, "cannot download file")
	}

	f, _, err := s.localFileService.LoadFile(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "cannot load file")
	}
	defer func() {
		//delete file as usage of local storage is only temporary
		s.localFileService.DeleteFile(fileName)
	}()

	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.New("cannot read from file")
	}

	_, sourceFormat, err := image.Decode(bytes.NewBuffer(fileBytes))
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode file")
	}

	if _, ok := s.availableFormats[Format(sourceFormat)]; !ok {
		return nil, errors.New(fmt.Sprintf("Unsupported format: %v", sourceFormat))
	}

	if cmd.SaveToCloud {
		return nil, s.cloudStorage.SaveFile(fileName, bytes.NewBuffer(fileBytes))
	}

	return fileBytes, nil
}

func sanitize(fileName string) string {
	fileName = path.Base(path.Clean(fileName))
	maxLen := maxFileNameLength
	if len(fileName) > maxLen {
		fileName = fileName[len(fileName)-maxLen:]
	}
	return fileName
}

func (s Service) DownloadFile(filePath string, url string) error {
	//todo - move to infrastructure
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = s.localFileService.SaveFile(filePath, resp.Body)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}
	return nil
}
