package test

import (
	"Gwallpaper"
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	res := Gwallpaper.CheckFolderHasImage(`D:\datacenter\壁纸\ForWallPaper\4Desktop`)
	println(res)
}

func TestGetPicPath(t *testing.T) {
	//Gwallpaper.C.FolderPath = `D:\datacenter\壁纸\ForWallPaper\4Desktop`
	path, err := Gwallpaper.C.GetPicPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(path)
}

/*func main() {
	folder := os.Args[1] // 获取命令行参数
	hasImage := false    // 标记是否有图片
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
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
				hasImage = true         // 设置标记为真
				fmt.Println(path)       // 打印图片路径
				return filepath.SkipDir // 跳过剩余的目录和文件
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	if hasImage {
		fmt.Println("该文件夹下有图片")
	} else {
		fmt.Println("该文件夹下没有图片")
	}
}*/
