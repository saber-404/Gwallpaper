package test

import (
	"Gwallpaper"
	"fmt"
	"strings"
	"testing"
)

func PrintTree(node *Gwallpaper.PicNode, depth int) {
	fmt.Printf("%s%s\n", strings.Repeat(" ", depth*4), node.Name)
	for _, child := range node.Children {
		PrintTree(child, depth+1)
	}
}

func TestInitSettings(t *testing.T) {
	Gwallpaper.InitSetting()
}

func TestEditConfig(t *testing.T) {
	Gwallpaper.EditConfig()
}

func TestUndoSetLockWallpaper(t *testing.T) {
	err := Gwallpaper.UndoSetLockWallpaper()
	if err != nil {
		Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
		return
	}
}

func TestConfig2Json(t *testing.T) {
	Gwallpaper.Config2Json(Gwallpaper.SleepTime, Gwallpaper.DefaultChangeLockWallPaper)
}

func TestChangeWallPaper(t *testing.T) {
	Gwallpaper.C.ChangeWallPaper()
}

func TestLoadData(t *testing.T) {
	PrintTree(&Gwallpaper.C.Cache, 0)
	Gwallpaper.C.Cache = Gwallpaper.PicNode{}
	Gwallpaper.LoadData()
	PrintTree(&Gwallpaper.C.Cache, 0)
}
