package utils

import (
	"path/filepath"
	"strings"
)

// SanitizeFilename - Очистка имени пользователя от двоеточий
func SanitizeFilename(name string) string {
	name = filepath.Base(name)
	name = strings.ReplaceAll(name, "..", "")
	return name
}

// SanitizeFolder - Очистка имени пользователя от двоеточий
func SanitizeFolder(folder string) string {
	folder = filepath.Base(folder)
	folder = strings.ReplaceAll(folder, "..", "")
	if folder == "" {
		folder = "default"
	}
	return folder
}
