package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

func GetStructFieldNames[T any]() []string {
	var names []string
	var t T
	tType := reflect.TypeOf(t)
	if tType.Kind() == reflect.Ptr {
		tType = tType.Elem()
	}
	for i := 0; i < tType.NumField(); i++ {
		names = append(names, tType.Field(i).Tag.Get("json"))
	}
	return names
}

func SaveUploadedFile(file *multipart.FileHeader, folder string) (string, error) {
	if file == nil {
		return "", nil
	}

	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create folder: %w", err)
	}

	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	fullPath := filepath.Join(folder, filename)

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return filepath.ToSlash(filepath.Join(folder, filename)), nil
}
