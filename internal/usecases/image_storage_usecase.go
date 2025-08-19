package usecases

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type ImageStorageUsecase interface {
	Save(fileHeader *multipart.FileHeader) (string, error)
}

type imageStorageUsecase struct {
	publicDir string
	baseURL   string
}

func NewImageStorageUsecase(publicDir string, baseURL string) ImageStorageUsecase {
	return &imageStorageUsecase{
		publicDir: publicDir,
		baseURL:   baseURL,
	}
}

// Save copies the uploaded file to the public directory and returns the public URL
func (u *imageStorageUsecase) Save(fileHeader *multipart.FileHeader) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	ext := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dstPath := filepath.Join(u.publicDir, filename)

	dst, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy file: %w", err)
	}

	return u.baseURL + filename, nil
}
