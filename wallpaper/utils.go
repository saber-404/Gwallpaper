package wallpaper

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func CheckFolderHasImage(folderpath string) bool {
	entries, err := os.ReadDir(folderpath)
	if err != nil {
		return false
	}
	for _, entry := range entries {
		path := filepath.Join(folderpath, entry.Name())
		if entry.IsDir() {
			if CheckFolderHasImage(path) {
				return true
			}
		} else if IsImage(path) {
			return true
		}
	}
	return false
}

func IsImage(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()
	_, _, err = image.DecodeConfig(file)
	if err != nil {
		return false
	}
	return true
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandIntn(length int) int {
	return rand.Intn(length)
}
