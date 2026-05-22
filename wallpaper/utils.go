package wallpaper

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func CheckFolderHasImage(folderpath string) bool {
	hasImage := false
	err := filepath.Walk(folderpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if IsImage(path) {
			hasImage = true
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return hasImage
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
