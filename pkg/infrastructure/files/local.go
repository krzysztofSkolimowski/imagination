package files

import (
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
)

const dirChmod = 0750

var (
	ErrUploadsDirEmpty = errors.New("uploadsDir cannot be empty")
	//todo - refactor errors
	ErrCannotSaveFile = errors.New("cannot save file")
)

type Validator interface{ Validate(r io.Reader) error }

type LocalFileService struct {
	uploadsDir UploadsDir
	validators []Validator
}

func (fs LocalFileService) CopyFile(srcFileID string, destFileID string, public bool) error {
	//todo - implement
	panic("implement me")
}

func NewLocalFileService(uploadsDir UploadsDir, validators []Validator) LocalFileService {
	if string(uploadsDir) == "" {
		panic(ErrUploadsDirEmpty)
	}

	ret := LocalFileService{
		uploadsDir,
		validators,
	}
	return ret
}

func (fs LocalFileService) SaveFile(fileID string, r io.Reader) error {
	fsPath, err := fs.SaveFromReader(fileID, r)
	if err != nil {
		return errors.Wrapf(ErrCannotSaveFile, "%s", fileID)
	}

	defer func() {
		if err != nil {
			removeErr := fs.DeleteFile(fileID)
			if removeErr != nil {
				panic(removeErr)
			}
		}
	}()

	file, err := os.Open(fsPath)
	defer func() {
		closeErr := file.Close()
		if err == nil && closeErr != nil {
			err = errors.Wrapf(closeErr, "could not close file %s", fileID)
		}
	}()

	if err != nil {
		return errors.Wrapf(err, "could not open file %s for validation", fileID)
	}
	for _, validator := range fs.validators {
		_, err = file.Seek(0, 0)
		if err != nil {
			return errors.Wrapf(err, "could not seek file %s", fileID)
		}
		err = validator.Validate(file)
		if err != nil {
			return errors.Wrapf(err, "could not validate file %s", fileID)
		}
	}

	return err
}

func (fs LocalFileService) SaveFromReader(filePath string, r io.Reader) (fsPath string, err error) {
	filePath = path.Join(string(fs.uploadsDir), filePath)
	dirpath := path.Dir(filePath)
	if err := createDir(dirpath); err != nil {
		return "", err
	}

	fsFile, err := os.Create(filePath)
	if err != nil {
		return "", errors.Wrap(err, "cannot create file")
	}

	if _, err := io.Copy(fsFile, r); err != nil {
		return "", errors.Wrap(err, "cannot save file")
	}

	return fsFile.Name(), nil
}

func createDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, dirChmod)
		if err != nil {
			return errors.Wrap(err, "cannot create file directory")
		}
	} else if err != nil {
		return errors.Wrap(err, "cannot check uploads dir existence")
	}

	return nil
}

func (fs LocalFileService) DeleteFile(fileID string) error {
	return os.Remove(path.Join(string(fs.uploadsDir), fileID))
}

func (fs LocalFileService) DeleteDirectory(directory string) error {
	return os.RemoveAll(path.Join(string(fs.uploadsDir), directory))
}

func (fs LocalFileService) LoadFile(fileID string) (io.ReadCloser, int, error) {
	f, err := os.Open(path.Join(string(fs.uploadsDir), fileID))
	if err != nil {
		return nil, 0, errors.Wrapf(err, "cannot open file")
	}

	// returning 0 to
	return f, 0, err
}
