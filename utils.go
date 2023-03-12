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

func RandIntn(length int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(length)
}

/*
// GetPicPath 随机给出图片路径
func (c *Config) GetPicPath() string {
	l := len(PicPath)
	return Prefix + PicPath[RandIntn(l)]
}

// GetPicPathSlice 获取图片路径字符串切片
func (c *Config) GetPicPathSlice() []string {
	var PathSlice []string
	err := filepath.Walk(C.FolderPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if IsImage(path) {
			PathSlice = append(PathSlice, path)
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return PathSlice
}

// SetPrefixAndPicPath 设定Prefix和PicPath切片
func (c *Config) SetPrefixAndPicPath() {
	PathSlice := c.GetPicPathSlice()
	l := len(PathSlice)
	if l == 0 {
		ShowMessage(errors.New("路径下没有图片"), MB_OK)
		return
	}
	prefix := PathSlice[0]
	s1 := make([]string, l)
	for _, str := range PathSlice {
		prefix = commonPrefix(prefix, str)
	}
	for i, str := range PathSlice {
		s1[i] = str[len(prefix):]
	}
	Prefix = prefix
	PicPath = s1
}

// commonPrefix returns the longest common prefix of two strings
func commonPrefix(a, b string) string {
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return a[:i]
}

*/

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

// GetIcon icon转byte流
/*func GetIcon(path string) (iconbytes []byte) {
	iconbytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return iconbytes
}*/

// IsDirEmpty 判断文件夹是否为空
/*func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}*/
