package files

import (
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateDirectory(t *testing.T) {
	cases := []struct {
		Name        string
		Directory   string
		ExpectedErr error
	}{
		{
			Name:        "empty",
			Directory:   "",
			ExpectedErr: ErrInvalidDirectory,
		},
		{
			Name:        "root",
			Directory:   "/",
			ExpectedErr: ErrInvalidDirectory,
		},
		{
			Name:        "missing_slash",
			Directory:   "/some-id",
			ExpectedErr: ErrInvalidDirectory,
		},
		{
			Name:        "missing_slash_subdir",
			Directory:   "/some-id/subdir",
			ExpectedErr: ErrInvalidDirectory,
		},
		{
			Name:        "valid",
			Directory:   "/some-id/",
			ExpectedErr: nil,
		},
		{
			Name:        "valid_subdir",
			Directory:   "/some-id/subdir/",
			ExpectedErr: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			err := validateDirectory(c.Directory)
			assert.Equal(t, c.ExpectedErr, err)
		})
	}
}

func TestProcessChunks(t *testing.T) {
	var result, input []string
	var chunks [][]string

	for i := 0; i <= 10; i++ {
		input = append(input, strconv.Itoa(i))
	}

	processor := func(chunk []string) error {
		result = append(result, chunk...)
		chunks = append(chunks, chunk)
		return nil
	}

	expected := make([]string, len(input))
	copy(expected, input)

	expectedChunks := [][]string{
		{"0", "1", "2"},
		{"3", "4", "5"},
		{"6", "7", "8"},
		{"9", "10"},
	}

	err := processChunks(input, 3, processor)
	require.NoError(t, err)
	require.ElementsMatch(t, expected, result)
	require.ElementsMatch(t, expectedChunks, chunks)
}
