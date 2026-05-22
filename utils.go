/*
工具函数
*/

package Gwallpaper

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

// CheckFolderHasImage 判断文件夹下是否有图片
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

// IsImage 判断是否是图片
func IsImage(path string) bool {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return false
	}
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


