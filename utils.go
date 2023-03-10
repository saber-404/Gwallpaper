/*
工具函数
*/

package Gwallpaper

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// GetPicName 获取图片名称
/*func (c *Config) GetPicName() (PicName string, err error) {
	files, err := os.ReadDir(c.FolderPath)
	if err != nil {
		return
	}
	rand.Seed(time.Now().UnixNano())
	retry := 0
	index := rand.Intn(len(files))
	for !isImage(files[index]) {
		rand.Seed(time.Now().UnixNano())
		index = rand.Intn(len(files))
		retry += 1
		if retry == c.RetryTimes {
			return "", errors.New("此文件夹下没有图片")
		}
	}
	return files[index].Name(), nil
}*/

// 判断是否是图片
/*func isImage(file fs.DirEntry) bool {
	if file.IsDir() {
		return false
	}
	//file.Info()
	ext := strings.ToLower(filepath.Ext(file.Name()))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}*/

// GetIcon icon转byte流
/*func GetIcon(path string) (iconbytes []byte) {
	iconbytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return iconbytes
}*/

// GetPicPath 随机给出图片路径
func (c *Config) GetPicPath() (PicPath string, err error) {
	var picpaths []string
	randompic := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, _, err = image.DecodeConfig(file)
			if err == nil {
				picpaths = append(picpaths, path)
			}
		}
		return nil
	}
	err = filepath.Walk(C.FolderPath, randompic)
	if err != nil {
		return "", err
	}
	l := len(picpaths)
	if l == 0 {
		return "", errors.New("此文件夹下没有图片")
	} else {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(l)
		return picpaths[index], nil
	}
}

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}
	return 0
}

// CheckFolderHasImage 判断文件夹下是否有图片
func CheckFolderHasImage(folderpath string) bool {
	hasImage := false
	err := filepath.Walk(folderpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() { // 如果不是目录
			file, err := os.Open(path) // 打开文件
			if err != nil {
				return err
			}
			defer file.Close()
			_, _, err = image.DecodeConfig(file) // 尝试解码图片
			if err == nil {                      // 如果没有错误
				hasImage = true // 设置标记为真
				//fmt.Println(path)       // 打印图片路径
				return filepath.SkipDir // 跳过剩余的目录和文件
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return hasImage
}
