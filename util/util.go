package util

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func CheckAndFileCreation(fileName string, body []byte) (*os.File, error) {
	_, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	err = file.Chmod(0777)
	if err != nil {
		return nil, err
	}

	_, err = file.Write(body)
	if err != nil {
		return nil, err
	}
	err = file.Sync()
	if err != nil {
		return nil, err
	}
	return file, nil

	// fmt.Printf("The content of %v was already downloaded, next. \n", fileName)
	return nil, nil
}

func SanitizeUrlFileName(url string) (fileName string) {
	fileName = fmt.Sprintf("Content_%s.txt", url)
	fileName = strings.ReplaceAll(fileName, "https://", "")
	fileName = strings.ReplaceAll(fileName, "/", "_")

	return fileName
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
