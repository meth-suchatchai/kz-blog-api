package kzfile

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"os"
)

func CreateTempDir() string {
	dirName := fmt.Sprintf("storage/%s", GenerateRandomTime())

	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		log.Errorf("error make temp directory: %v", err)
		return ""
	}

	return dirName
}
