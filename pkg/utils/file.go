package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func UploadImage(uploadDir string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
			return "", err
		}
	}

	timestamp := time.Now().UnixNano()
	extension := filepath.Ext(file.Filename)
	uniqueFilename := fmt.Sprintf("%d%s", timestamp, extension)

	filePath := filepath.Join(uploadDir, uniqueFilename)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return filePath, nil
}

func DeleteImage(path string) error {
	filePath := filepath.Join(path)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return err
	}

	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

func GenerateUniqueFileName(originalFileName string) string {
	sanitizedFileName := strings.ReplaceAll(originalFileName, " ", "-")

	ext := filepath.Ext(sanitizedFileName)

	baseName := strings.TrimSuffix(originalFileName, ext)

	timestamp := time.Now().Unix()

	uniqueFileName := fmt.Sprintf("%d_%s%s", timestamp, baseName, ext)

	return uniqueFileName
}
