package files

import (
	"path"
)

type PathResolver struct {
	uploadsDir string
}

func NewPathResolver(uploadsDir string) (PathResolver, error) {
	if uploadsDir == "" {
		return PathResolver{}, ErrUploadsDirEmpty
	}

	return PathResolver{uploadsDir}, nil
}

func (p PathResolver) Resolve(fileID string) string {
	return path.Join(p.uploadsDir, fileID)
}
