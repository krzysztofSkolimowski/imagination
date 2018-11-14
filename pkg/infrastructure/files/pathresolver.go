package files

import (
	"path"
)

type pathResolver struct {
	uploadsDir string
}

func NewPathResolver(uploadsDir string) (pathResolver, error) {
	if uploadsDir == "" {
		return pathResolver{}, ErrUploadsDirEmpty
	}

	return pathResolver{uploadsDir}, nil
}

func (p pathResolver) Resolve(fileID string) string {
	return path.Join(p.uploadsDir, fileID)
}
