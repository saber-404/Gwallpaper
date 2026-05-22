package test

import (
	"Gwallpaper"
	"fmt"
	"testing"
	"time"
)

func TestSetTreeNode(t *testing.T) {
	Gwallpaper.SetTreeNode()
	PrintTree(&Gwallpaper.C.Cache, 0)
	Gwallpaper.C.Cache.Name = `D:\datacenter\壁纸\ForWallPaper\TryTree`
	Gwallpaper.SetTreeNode()
	PrintTree(&Gwallpaper.C.Cache, 0)
}

func TestGetPicPathByTree(t *testing.T) {
	for i := 0; i < 10; i++ {
		picpath := Gwallpaper.C.GetPicPathByTree()
		fmt.Println(picpath)
		time.Sleep(2 * time.Second)
	}
}


