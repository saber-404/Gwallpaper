package main

import (
	"Gwallpaper"
	"time"
)

func ChangeIt() {
	PicName, err := Gwallpaper.C.GetPicName()
	if err != nil {
		Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
		return
	}
	err = Gwallpaper.SetWallpaper(Gwallpaper.C.FolderPath + PicName)
	if err != nil {
		Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
		return
	}
}

func main() {
	for {
		select {
		case <-time.After(time.Duration(Gwallpaper.C.SleepTime) * time.Second):
			// 更换壁纸的代码逻辑
			ChangeIt()
		}
	}
}
