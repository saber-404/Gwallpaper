package main

import (
	"Gwallpaper"
	"github.com/kardianos/service"
	"os"
	"time"
)

var svcConfig = &service.Config{
	Name:        "AWallpaperChanger",
	DisplayName: "AWallpaperChanger",
	Description: "Automatically change wallpaper every 15 minutes",
}

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

type ChangeWallPaperService struct {
	stopChan chan struct{}
}

func (s *ChangeWallPaperService) Start(svc service.Service) error {
	// 初始化服务配置
	s.stopChan = make(chan struct{})

	// 在新的 goroutine 中运行您的程序逻辑
	go func() {
		for {
			select {
			case <-s.stopChan:
				return
			case <-time.After(time.Duration(Gwallpaper.C.SleepTime) * time.Second):
				// 更换壁纸的代码逻辑
				ChangeIt()
			}
		}
	}()

	return nil
}
func (s *ChangeWallPaperService) Stop(svc service.Service) error {
	// 停止您的程序逻辑
	close(s.stopChan)
	return nil
}

func main() {
	s := &ChangeWallPaperService{}
	svc, err := service.New(s, svcConfig)
	if err != nil {
		return
	}

	if len(os.Args) == 1 {
		svc.Run()
		return
	}
	cmd := os.Args[1]
	if cmd == "install" {
		err := svc.Install()
		if err != nil {
			Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
			return
		}
	}
	if cmd == "uninstall" {
		err := svc.Uninstall()
		if err != nil {
			Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
			return
		}
	}
	if cmd == "start" {
		err := svc.Start()
		if err != nil {
			Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
			return
		}
	}
	if cmd == "stop" {

		err := svc.Stop()
		if err != nil {
			Gwallpaper.ShowMessage(err, Gwallpaper.MB_OK)
			return
		}
	}
}

/*package main

import "Gwallpaper"

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
	ChangeIt()
}*/
