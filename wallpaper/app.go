package wallpaper

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type App struct {
	Config       Config
	picTree      PicNode
	lastLockPath string
}

func New() *App {
	a := &App{}
	a.InitSetting()
	return a
}

func (a *App) ChangeWallPaper() {
	path := a.GetPicPathByTree()
	if !IsImage(path) {
		a.SetTreeNode()
	}
	err := SetWallpaper(path)
	if err != nil {
		ShowMessage(err, MB_OK)
		return
	}
	if a.Config.ChangeLockWallPaper && path != a.lastLockPath {
		err := setLockWallpaper(path)
		if err != nil {
			ShowMessage(err, MB_OK)
			return
		}
		a.lastLockPath = path
	}
}

func (a *App) ResetDefaults() {
	a.Config.SleepTime = SleepTime
	a.Config.ChangeLockWallPaper = DefaultChangeLockWallPaper
	a.lastLockPath = ""
	a.SaveData()
}

func (a *App) InitSetting() {
	_, err := os.Stat("setting.json")
	if err != nil {
		a.Config2Json(SleepTime, DefaultChangeLockWallPaper)
		return
	}
	a.LoadData()
	a.SetTreeNode()
	if !CheckFolderHasImage(a.Config.FolderPath) {
		ShowMessage(errors.New("壁纸文件夹内无图片"), MB_OK)
		a.Config2Json(a.Config.SleepTime, a.Config.ChangeLockWallPaper)
	}
}

func (a *App) Config2Json(sleepTime int64, changeLock bool) {
	IsChoice, PicFolderPath := ShowFolderDialogForGetFolderPath("选择壁纸文件夹")
	if !IsChoice {
		os.Exit(0)
	}
	hasImage := CheckFolderHasImage(PicFolderPath)
	if !hasImage {
		IsChoice, PicFolderPath = ShowFolderDialogForGetFolderPath("前面的文件夹下没有图片，请重新选择")
	}
	if !IsChoice {
		os.Exit(0)
	}
	a.Config.SleepTime = sleepTime
	a.Config.ChangeLockWallPaper = changeLock
	a.Config.FolderPath = PicFolderPath
	a.SetTreeNode()
	a.SaveData()
}

func (a *App) SaveData() error {
	bytes, err := json.MarshalIndent(a.Config, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile("./setting.json", bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) LoadData() {
	file, err := os.ReadFile("setting.json")
	if err != nil {
		ShowMessage(errors.New("创建默认setting.json失败"), MB_OK)
		os.Exit(0)
		return
	}
	err = json.Unmarshal(file, &a.Config)
	if err != nil {
		ShowMessage(errors.New("json文件解析失败"), MB_OK)
		os.Exit(1)
		return
	}
}

func (a *App) SetTreeNode() {
	a.picTree.Children = nil
	a.picTree.Insert(a.Config.FolderPath)
}

func (a *App) GetPicPathByTree() string {
	if len(a.picTree.Children) == 0 {
		ShowMessage(errors.New("壁纸文件夹内无图片"), MB_OK)
		a.Config2Json(a.Config.SleepTime, a.Config.ChangeLockWallPaper)
	}
	var path string
	node := a.picTree
	for {
		l := len(node.Children)
		if l == 0 {
			break
		}
		child := node.Children[RandIntn(l)]
		path = filepath.Join(path, child.Name)
		node = *child
	}
	return filepath.Join(a.Config.FolderPath, path)
}
