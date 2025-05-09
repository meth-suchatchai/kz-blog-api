package fileservices

import (
	"github.com/gofiber/fiber/v2"
	filemodels "github.com/meth-suchatchai/kz-blog-api/app/file/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	"github.com/meth-suchatchai/kz-blog-api/lib/kzfile"
	"io"
	"os"

	"mime/multipart"
)

func (svc *defaultService) FileUpload(multipart *multipart.FileHeader, dir ...string) (string, *fiber.Error) {
	tempDir := kzfile.CreateTempDir()
	file, err := multipart.Open()
	if err != nil {
		return "", errors.NewDefaultFiberError(err)
	}
	defer file.Close()

	contentType := ""
	directory := ""
	if len(multipart.Header["Content-Type"]) > 0 && multipart.Header["Content-Type"][0] != "" {
		contentType = multipart.Header["Content-Type"][0]
	}
	if len(dir) > 0 {
		directory = dir[0]
	}

	tmpFile, err := kzfile.CreateTempFile(tempDir, multipart.Filename)
	if err != nil {
		return "", errors.NewDefaultFiberError(err)
	}
	defer tmpFile.Close()

	/* Copy Uploaded File */
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		return "", errors.NewDefaultFiberError(err)
	}

	path, vErr := svc.fileRepo.UploadFile(&filemodels.UploadFileData{
		Filename:  multipart.Filename,
		Directory: directory,
		MimeType:  contentType,
		Size:      multipart.Size,
		File:      tmpFile,
	})
	if vErr != nil {
		return "", errors.NewDefaultFiberError(err)
	}

	/* Delete Folder */
	defer os.RemoveAll(tempDir)

	return path, nil
}
