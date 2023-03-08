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
	exitItem := systray.AddMenuItem("退出程序", "Exit app")
	reloadItem := systray.AddMenuItem("重载配置", "Reload setting")
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
			}
		}
	}()
}
func onExit() {

}
