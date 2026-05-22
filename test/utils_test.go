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

func TestIsImage(t *testing.T) {
	fmt.Println(Gwallpaper.IsImage(`D:\datacenter\壁纸\ForWallPaper\TryTree\1.txt`))
	fmt.Println(Gwallpaper.IsImage(`D:\datacenter\壁纸\ForWallPaper\TryTree\subTryTree1.jpg`))
	fmt.Println(Gwallpaper.IsImage(`D:\datacenter\壁纸\ForWallPaper\TryTree\`))
}


