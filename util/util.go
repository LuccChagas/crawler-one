package util

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CheckAndFileCreation(fileName string, body []byte) *os.File {
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			// TODO: create a nice Log message with error details (Try to put in a recovery logic)
			log.Fatal(err)
		}

		err = file.Chmod(0777)
		if err != nil {
			log.Fatal(err)
		}

		_, err = file.Write(body)
		if err != nil {
			log.Fatal(err)
		}
		file.Sync()
		return file
	}
	fmt.Printf("The content of %v was already downloaded, next.", fileName)
	return nil
}

func SanitizeUrlFileName(url string) (fileName string) {
	fileName = fmt.Sprintf("Content_%s.txt", url)
	fileName = strings.ReplaceAll(fileName, "https://", "")
	fileName = strings.ReplaceAll(fileName, "/", "_")

	return fileName
}
