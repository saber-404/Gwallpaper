package test

import (
	"Gwallpaper"
	"errors"
	"testing"
)

func TestSetWallpaper(t *testing.T) {
	newWallpaperPath := "D:\\test.png"
	err := Gwallpaper.SetWallpaper(newWallpaperPath)
	if err != nil {
		t.Log(err)
		return
	}
}

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
