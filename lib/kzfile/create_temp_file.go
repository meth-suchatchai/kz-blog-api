package kzfile

import (
	"fmt"
	"os"
)

func CreateTempFile(dir, file string) (*os.File, error) {
	tempFile := fmt.Sprintf("%s/%s", dir, file)
	f, err := os.Create(tempFile)
	if err != nil {
		return nil, err
	}

	return f, nil
}
