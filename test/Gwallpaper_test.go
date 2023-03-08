package test

import (
	"Gwallpaper"
	"errors"
	"testing"
)

//func TestSetWallpaper(t *testing.T) {
//	newWallpaperPath := "D:\\test.png"
//	err := Gwallpaper.SetWallpaper(newWallpaperPath)
//	if err != nil {
//		t.Log(err)
//		return
//	}
//}

func TestGetPicName(t *testing.T) {
	name, err := Gwallpaper.C.GetPicName()
	if err != nil {
		t.Log(err)
		return
	}
	t.Fatal(name)
}

func TestShowMessage(t *testing.T) {
	err_test := errors.New("test error")
	Gwallpaper.ShowMessage(err_test, Gwallpaper.MB_OK)
}

//func TestSetLockWallpaper(t *testing.T) {
//	err := Gwallpaper.SetLockWallpaper("D:\\datacenter\\壁纸\\ForWallPaper\\4Browser\\105956477_p0.png")
//	if err != nil {
//		Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
//		return
//	}
//}

//func TestUndoSetLockWallpaper(t *testing.T) {
//	err := Gwallpaper.UndoSetLockWallpaper()
//	if err != nil {
//		Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
//		return
//	}
//}
