package main

import (
	"Gwallpaper"
	"github.com/getlantern/systray"
	"time"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(Gwallpaper.Icon)
	systray.SetTitle(Gwallpaper.Title)
	systray.SetTooltip(Gwallpaper.Title)
	reloadItem := systray.AddMenuItem("应用配置", "Reload setting")
	defaultItem := systray.AddMenuItem("恢复默认", "reset settings")
	editItem := systray.AddMenuItem("编辑配置", "Edit Config File")
	changeItem := systray.AddMenuItem("换一张", "Choose other")
	LockItem := systray.AddMenuItemCheckbox("改变锁屏", "test1", Gwallpaper.C.ChangLockWallPaper)
	systray.AddSeparator()
	exitItem := systray.AddMenuItem("退出程序", "Exit app")
	go func() {
		for {
			select {
			//退出
			case <-exitItem.ClickedCh:
				systray.Quit()
				return
			//	定时更换壁纸
			case <-time.After(time.Duration(Gwallpaper.C.SleepTime) * time.Second):
				Gwallpaper.C.ChangeWallPaper()
			//	应用配置文件
			case <-reloadItem.ClickedCh:
				Gwallpaper.LoadData()
				Gwallpaper.SetTreeNode()
				Gwallpaper.C.ChangeWallPaper()
				Gwallpaper.SaveData(Gwallpaper.C)
			//	换一张壁纸
			case <-changeItem.ClickedCh:
				Gwallpaper.C.ChangeWallPaper()
			//	恢复默认配置并立即应用
			case <-defaultItem.ClickedCh:
				Gwallpaper.Config2Json(Gwallpaper.SleepTime, Gwallpaper.ChangeLockWallPaper)
				Gwallpaper.InitSetting()
				Gwallpaper.C.ChangeWallPaper()
			//	改变锁屏
			case <-LockItem.ClickedCh:
				Gwallpaper.C.ChangLockWallPaper = !Gwallpaper.C.ChangLockWallPaper
				if Gwallpaper.C.ChangLockWallPaper {
					LockItem.Check()
					Gwallpaper.C.ChangeWallPaper()
				} else {
					LockItem.Uncheck()
					err := Gwallpaper.UndoSetLockWallpaper()
					if err != nil {
						Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
					}
				}
			//	编辑配置
			case <-editItem.ClickedCh:
				Gwallpaper.EditConfig()
			}
		}
	}()
}
func onExit() {
	Gwallpaper.SaveData(Gwallpaper.C)
}
