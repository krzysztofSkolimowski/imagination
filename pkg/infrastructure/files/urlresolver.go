package files

import (
	"github.com/pkg/errors"
	"net/url"
	"path"
)

var ErrFailedToParseURL = errors.New("Failed to parse URL")

type URLResolver struct {
	baseURL *url.URL
}

func NewURLResolver(baseURL string) (URLResolver, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return URLResolver{}, errors.Wrapf(ErrFailedToParseURL, baseURL)
	}
	return URLResolver{baseURL: parsedURL}, nil
}

func (u URLResolver) Resolve(fileID string) string {
	resolved := *u.baseURL
	resolved.Path = path.Join(resolved.Path, fileID)
	return resolved.String()
}
