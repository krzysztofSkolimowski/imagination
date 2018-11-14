package files

import "io"

type FileService interface {
	SaveFile(fileID string, r io.Reader) error
	LoadFile(fileID string) (io.ReadCloser, int, error)
	DeleteFile(fileID string) error
	DeleteDirectory(directory string) error
	CopyFile(srcFileID string, destFileID string, public bool) error
}

type AccessControlFileService interface {
	FileService
	SavePrivateFile(fileID string, r io.Reader) error
}

type Validator interface {
	Validate(r io.Reader) error
}
