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
	//设置托盘图标和提示文本
	systray.SetIcon(Gwallpaper.GetIcon("icon.ico"))
	systray.SetTitle(Gwallpaper.Title)
	systray.SetTooltip(Gwallpaper.Title)
	reloadItem := systray.AddMenuItem("重载配置", "Reload setting")
	changeItem := systray.AddMenuItem("换一张", "Choose other")
	LockItem := systray.AddMenuItemCheckbox("改变锁屏", "test1", Gwallpaper.C.ChangLockWallPaper)
	systray.AddSeparator()
	exitItem := systray.AddMenuItem("退出程序", "Exit app")
	go func() {
		for {
			select {
			case <-exitItem.ClickedCh:
				systray.Quit()
				return
			case <-time.After(time.Duration(Gwallpaper.C.SleepTime) * time.Second):
				// 更换壁纸的代码逻辑
				Gwallpaper.C.ChangeWallPaper()
			case <-reloadItem.ClickedCh:
				Gwallpaper.InitSetting()
			case <-changeItem.ClickedCh:
				Gwallpaper.C.ChangeWallPaper()
			case <-LockItem.ClickedCh:
				Gwallpaper.C.ChangLockWallPaper = !Gwallpaper.C.ChangLockWallPaper
				//fmt.Println("点击动作")
				if Gwallpaper.C.ChangLockWallPaper {
					LockItem.Check()
					//fmt.Println("执行改变锁屏")
				} else {
					LockItem.Uncheck()
					//fmt.Println("执行恢复锁屏")
					err := Gwallpaper.UndoSetLockWallpaper()
					if err != nil {
						Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
						return
					}
				}
			}
		}
	}()
}
func onExit() {

}
