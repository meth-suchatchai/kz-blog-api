package filemodels

import (
	"os"
)

type UploadFileData struct {
	Filename  string
	Directory string
	MimeType  string
	Size      int64
	File      *os.File
}
