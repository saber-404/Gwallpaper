package test

import (
	"Gwallpaper"
	"testing"
)

//func TestGetPicName(t *testing.T) {
//	name, err := Gwallpaper.C.GetPicName()
//	if err != nil {
//		t.Log(err)
//		return
//	}
//	t.Fatal(name)
//}

func TestInitSettings(t *testing.T) {
	Gwallpaper.InitSetting()
}

func TestEditConfig(t *testing.T) {
	Gwallpaper.EditConfig()
}

//func TestSetLockWallpaper(t *testing.T) {
//	err := Gwallpaper.SetLockWallpaper("D:\\datacenter\\壁纸\\ForWallPaper\\4Browser\\105956477_p0.png")
//	if err != nil {
//		Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
//		return
//	}
//}

func TestUndoSetLockWallpaper(t *testing.T) {
	err := Gwallpaper.UndoSetLockWallpaper()
	if err != nil {
		Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
		return
	}
}

func TestConfig2Json(t *testing.T) {
	Gwallpaper.Config2Json()
}

func TestChangeWallPaper(t *testing.T) {
	Gwallpaper.C.ChangeWallPaper()
}
