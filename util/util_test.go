package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckAndFileCreation(t *testing.T) {
	fileName := RandomString(6)
	body := make([]byte, 1)

	xfile, err := os.Create(fileName)
	require.NoError(t, err)

	file, err := CheckAndFileCreation(fileName, body)
	if os.IsNotExist(err) {
		require.NotNil(t, file)
		require.NoError(t, err)
		require.NotEmpty(t, file)
	}

	file, err = CheckAndFileCreation("random_file", body)
	if os.IsNotExist(err) {
		require.NotNil(t, file)
		require.NoError(t, err)
		require.NotEmpty(t, file)
	}

	require.NoError(t, err)
	require.IsType(t, xfile, file)
}

func TestSanitizeUrlFileName(t *testing.T) {
	url := "www.google.com"

	filename := SanitizeUrlFileName(url)
	require.NotEmpty(t, filename)

}
