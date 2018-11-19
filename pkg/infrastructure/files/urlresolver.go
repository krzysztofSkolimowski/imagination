package files

import (
	"github.com/pkg/errors"
	"net/url"
	"path"
)

var ErrFailedToParseURL = errors.New("Failed to parse URL")

type BaseURL string

type URLResolver struct {
	baseURL *url.URL
}

func NewURLResolver(baseURL BaseURL) (URLResolver, error) {
	parsedURL, err := url.Parse(string(baseURL))
	if err != nil {
		return URLResolver{}, errors.Wrapf(ErrFailedToParseURL, string(baseURL))
	}
	return URLResolver{baseURL: parsedURL}, nil
}

func (u URLResolver) Resolve(fileID string) string {
	resolved := *u.baseURL
	resolved.Path = path.Join(resolved.Path, fileID)
	return resolved.String()
}
