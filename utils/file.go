package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func EnsureDir(path string) {
	err := os.MkdirAll(path, os.ModePerm) // os.ModePerm = 0777
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}
}

func GenerateFilename(ext string) string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%d%s", timestamp, ext)
}
