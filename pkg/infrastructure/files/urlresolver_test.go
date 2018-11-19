package files_test

import (
	"github.com/krzysztofSkolimowski/imagination/pkg/infrastructure/files"
	"github.com/pkg/errors"
	"testing"
)

func TestURLResolverValid(t *testing.T) {
	cases := []struct {
		Name          string
		BaseURL       string
		ID            string
		Expected      string
		ExpectedError error
	}{
		{
			Name:     "valid url and id",
			BaseURL:  "https://example.test",
			ID:       "any-valid-id",
			Expected: "https://example.test/any-valid-id",
		},
		{
			Name:     "redundant slashes",
			BaseURL:  "https://example.test///",
			ID:       "any-valid-id",
			Expected: "https://example.test/any-valid-id",
		},
		{
			Name:     "sub path in url",
			BaseURL:  "https://example.test/images/",
			ID:       "any-valid-id",
			Expected: "https://example.test/images/any-valid-id",
		},
		{
			Name:     "sub path in id",
			BaseURL:  "https://example.test",
			ID:       "/images/any-valid-id",
			Expected: "https://example.test/images/any-valid-id",
		},
		{
			Name:     "slashes in id",
			BaseURL:  "https://example.test",
			ID:       "//any-valid-id//",
			Expected: "https://example.test/any-valid-id",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			resolver, err := files.NewURLResolver(files.BaseURL(c.BaseURL))
			if errors.Cause(c.ExpectedError) != c.ExpectedError {
				t.Fatalf("Expected: %v, got %v", c.ExpectedError, err)
			}

			result := resolver.Resolve(c.ID)
			if c.Expected != result {
				t.Fatalf("Expected %v, got %v", c.Expected, result)
			}
		})
	}
}
