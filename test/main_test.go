package test

import (
	"Gwallpaper"
	"testing"
)

func TestReload(t *testing.T) {
	Gwallpaper.LoadData()
	Gwallpaper.SetTreeNode()
	Gwallpaper.C.ChangeWallPaper()
	Gwallpaper.SaveData(Gwallpaper.C)
}


