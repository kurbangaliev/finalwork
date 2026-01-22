package utils

import (
	"path/filepath"
	"strings"
)

func SanitizeFilename(name string) string {
	name = filepath.Base(name)
	name = strings.ReplaceAll(name, "..", "")
	return name
}

func SanitizeFolder(folder string) string {
	folder = filepath.Base(folder)
	folder = strings.ReplaceAll(folder, "..", "")
	if folder == "" {
		folder = "default"
	}
	return folder
}
