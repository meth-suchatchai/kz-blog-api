package main

import (
	"github.com/kuroshibaz/lib/kzfile"
	"log"
)

func main() {
	tempDir := kzfile.CreateTempDir()
	file, err := kzfile.CreateTempFile(tempDir, "./storage/bg-grid.png")
	log.Print(file, err)
}
