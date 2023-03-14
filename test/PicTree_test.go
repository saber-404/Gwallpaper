package test

import (
	"Gwallpaper"
	"fmt"
	"testing"
	"time"
)

func TestSetTreeNode(t *testing.T) {
	Gwallpaper.SetTreeNode()
	Gwallpaper.PrintTree(&Gwallpaper.C.Cache, 0)
	Gwallpaper.C.Cache.Name = `D:\datacenter\壁纸\ForWallPaper\TryTree`
	Gwallpaper.SetTreeNode()
	Gwallpaper.PrintTree(&Gwallpaper.C.Cache, 0)
}

func TestGetPicPathByTree(t *testing.T) {
	for i := 0; i < 10; i++ {
		picpath := Gwallpaper.C.GetPicPathByTree()
		fmt.Println(picpath)
		time.Sleep(2 * time.Second)
	}
}

/*func TestSaveData2File(t *testing.T) {
	Gwallpaper.SetTreeNode(0)
	Gwallpaper.C.FolderPath = `D:\datacenter\壁纸\ForWallPaper\TryTree`
	Gwallpaper.SetTreeNode(1)
	Gwallpaper.TreeNode.SaveData2File("./cache")
}*/

/*func TestLoadDataFromFile(t *testing.T) {
	Gwallpaper.TreeNode = Gwallpaper.PicNode{}
	Gwallpaper.PrintTree(&Gwallpaper.TreeNode, 0)
	Gwallpaper.TreeNode.LoadDataFromFile("./1.dat")
	Gwallpaper.PrintTree(&Gwallpaper.TreeNode, 0)
}*/
