package main

import (
	"Gwallpaper/wallpaper"
	"github.com/getlantern/systray"
	"time"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(wallpaper.Icon)
	systray.SetTitle(wallpaper.Title)
	systray.SetTooltip(wallpaper.Title)
	reloadItem := systray.AddMenuItem("应用配置", "Reload setting")
	defaultItem := systray.AddMenuItem("恢复默认", "reset settings")
	editItem := systray.AddMenuItem("编辑配置", "Edit Config File")
	changeItem := systray.AddMenuItem("换一张", "Choose other")
	LockItem := systray.AddMenuItemCheckbox("改变锁屏", "test1", wallpaper.C.ChangeLockWallPaper)
	systray.AddSeparator()
	exitItem := systray.AddMenuItem("退出程序", "Exit app")
	go func() {
		for {
			select {
			case <-exitItem.ClickedCh:
				systray.Quit()
				return
			case <-time.After(time.Duration(wallpaper.C.SleepTime) * time.Second):
				wallpaper.C.ChangeWallPaper()
			case <-reloadItem.ClickedCh:
				wallpaper.LoadData()
				wallpaper.SetTreeNode()
				wallpaper.C.ChangeWallPaper()
				wallpaper.SaveData(wallpaper.C)
			case <-changeItem.ClickedCh:
				wallpaper.C.ChangeWallPaper()
			case <-defaultItem.ClickedCh:
				wallpaper.Config2Json(wallpaper.SleepTime, wallpaper.DefaultChangeLockWallPaper)
				wallpaper.InitSetting()
				wallpaper.C.ChangeWallPaper()
			case <-LockItem.ClickedCh:
				wallpaper.C.ChangeLockWallPaper = !wallpaper.C.ChangeLockWallPaper
				if wallpaper.C.ChangeLockWallPaper {
					LockItem.Check()
					wallpaper.C.ChangeWallPaper()
				} else {
					LockItem.Uncheck()
					err := wallpaper.UndoSetLockWallpaper()
					if err != nil {
						wallpaper.ShowMessage(err, wallpaper.MB_OK)
					}
				}
			case <-editItem.ClickedCh:
				wallpaper.EditConfig()
			}
		}
	}()
}

func onExit() {
	wallpaper.SaveData(wallpaper.C)
}
