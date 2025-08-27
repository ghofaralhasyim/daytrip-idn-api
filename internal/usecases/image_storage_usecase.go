package usecases

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type ImageStorageUsecase interface {
	Save(path string, fileHeader *multipart.FileHeader) (string, error)
	Delete(filePath string) error
}

type imageStorageUsecase struct {
	publicDir string
}

func NewImageStorageUsecase(publicDir string) ImageStorageUsecase {
	return &imageStorageUsecase{
		publicDir: publicDir,
	}
}

// Save copies the uploaded file to the public directory and returns the public URL
func (u *imageStorageUsecase) Save(path string, fileHeader *multipart.FileHeader) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	ext := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dstPath := filepath.Join(u.publicDir, path, filename)

	log.Println("\n ================> saving image", dstPath)

	dst, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy file: %w", err)
	}

	return fmt.Sprintf("%s%s", path, filename), nil
}

func (u *imageStorageUsecase) Delete(filePath string) error {
	log.Println("\n ================> Deleting file:", filePath)

	// Build full path if the filePath is relative
	fullPath := filePath
	if !filepath.IsAbs(filePath) {
		fullPath = filepath.Join(u.publicDir, filePath)
	}

	// Check if the file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", fullPath)
	}

	// Remove the file
	if err := os.Remove(fullPath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}
