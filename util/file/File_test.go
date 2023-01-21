package file

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestValidateFileType(t *testing.T) {
	file, err := os.Open("images.png")
	require.NoError(t, err)
	err = ValidateFileType(file)
	require.NoError(t, err)
}

func TestErrorValidateFileType(t *testing.T) {
	file, err := os.Open("samplepptx.pptx")
	require.NoError(t, err)
	err = ValidateFileType(file)
	require.Error(t, err, "invalid type! must be [image/jpg image/jpeg image/gif image/png image/webp video/mp4]")
}
