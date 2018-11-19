package files

import (
	"path"
)

type UploadsDir string

type PathResolver struct {
	uploadsDir UploadsDir
}

func NewPathResolver(u UploadsDir) (PathResolver, error) {
	if string(u) == "" {
		return PathResolver{}, ErrUploadsDirEmpty
	}

	return PathResolver{u}, nil
}

func (p PathResolver) Resolve(fileID string) string {
	return path.Join(string(p.uploadsDir), fileID)
}
