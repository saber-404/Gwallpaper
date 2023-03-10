package test

import (
	"Gwallpaper"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"testing"
)

func TestCheckFolderHasImage(t *testing.T) {
	res := Gwallpaper.CheckFolderHasImage(`D:\datacenter\壁纸\ForWallPaper\4Desktop`)
	println(res)
}

func TestGetPicPath(t *testing.T) {
	path := Gwallpaper.C.GetPicPath()
	fmt.Println(path)
}

func TestIsImage(t *testing.T) {
	fmt.Println(Gwallpaper.IsImage(`D:\datacenter\壁纸\ForWallPaper\TryTree\1.txt`))
	fmt.Println(Gwallpaper.IsImage(`D:\datacenter\壁纸\ForWallPaper\TryTree\subTryTree1.jpg`))
	fmt.Println(Gwallpaper.IsImage(`D:\datacenter\壁纸\ForWallPaper\TryTree\`))
}

func TestSetPrefixAndPicPath(t *testing.T) {
	// 由于SetPrefixAndPicPath调用了GetPicPath GetPicPath可以不用测试
	Gwallpaper.C.SetPrefixAndPicPath()
	fmt.Println(Gwallpaper.Prefix)
	fmt.Println(Gwallpaper.PicPath)

	Gwallpaper.C.FolderPath = `D:\datacenter\壁纸\ForWallPaper\TryTree`
	Gwallpaper.C.SetPrefixAndPicPath()
	fmt.Println(Gwallpaper.Prefix)
	fmt.Println(Gwallpaper.PicPath)
}
