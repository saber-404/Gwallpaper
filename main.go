package main

import (
	"Gwallpaper/wallpaper"
	"time"

	"github.com/getlantern/systray"
)

var app *wallpaper.App

func main() {
	app = wallpaper.New()
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
	LockItem := systray.AddMenuItemCheckbox("改变锁屏", "test1", app.Config.ChangeLockWallPaper)
	systray.AddSeparator()
	exitItem := systray.AddMenuItem("退出程序", "Exit app")
	go func() {
		for {
			select {
			case <-exitItem.ClickedCh:
				systray.Quit()
				return
			case <-time.After(time.Duration(app.Config.SleepTime) * time.Second):
				app.ChangeWallPaper()
			case <-reloadItem.ClickedCh:
				app.LoadData()
				if app.Config.ChangeLockWallPaper {
					LockItem.Check()
				} else {
					LockItem.Uncheck()
					err := wallpaper.UndoSetLockWallpaper()
					if err != nil {
						wallpaper.ShowMessage(err, wallpaper.MB_OK)
					}
				}
				app.SetTreeNode()
				app.ChangeWallPaper()
				app.SaveData()
			case <-changeItem.ClickedCh:
				app.ChangeWallPaper()
			case <-defaultItem.ClickedCh:
				app.ResetDefaults()
				LockItem.Uncheck()
				if err := wallpaper.UndoSetLockWallpaper(); err != nil {
					wallpaper.ShowMessage(err, wallpaper.MB_OK)
				}
				app.ChangeWallPaper()
			case <-LockItem.ClickedCh:
				app.Config.ChangeLockWallPaper = !app.Config.ChangeLockWallPaper
				if app.Config.ChangeLockWallPaper {
					LockItem.Check()
					app.ChangeWallPaper()
				} else {
					LockItem.Uncheck()
					err := wallpaper.UndoSetLockWallpaper()
					if err != nil {
						wallpaper.ShowMessage(err, wallpaper.MB_OK)
					}
				}
				app.SaveData()
			case <-editItem.ClickedCh:
				wallpaper.EditConfig()
			}
		}
	}()
}

func onExit() {
	app.SaveData()
}
