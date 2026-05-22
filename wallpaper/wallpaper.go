package wallpaper

import (
	_ "embed"
	"encoding/json"
	"errors"
	"golang.org/x/sys/windows/registry"
	"os"
	"os/exec"
)

var (
	C       Config
	picTree PicNode

	//go:embed icon.ico
	Icon []byte
)

const (
	Title                      = "GwallPaper"
	LockWallPaperRegPath       = `SOFTWARE\Microsoft\Windows\CurrentVersion\PersonalizationCSP`
	SleepTime            int64 = 900
	DefaultChangeLockWallPaper = false
)

type Config struct {
	SleepTime          int64
	ChangeLockWallPaper bool
	FolderPath         string
}

func init() {
	InitSetting()
}

func (c *Config) ChangeWallPaper() {
	path := GetPicPathByTree()
	if !IsImage(path) {
		SetTreeNode()
	}
	err := SetWallpaper(path)
	if err != nil {
		ShowMessage(err, MB_OK)
		return
	}
	if c.ChangeLockWallPaper {
		err := setLockWallpaper(path)
		if err != nil {
			ShowMessage(err, MB_OK)
			return
		}
	}
}

func InitSetting() {
	_, err := os.Stat("setting.json")
	if err != nil {
		Config2Json(SleepTime, DefaultChangeLockWallPaper)
		return
	}
	LoadData()
	if !CheckFolderHasImage(C.FolderPath) {
		ShowMessage(errors.New("壁纸文件夹内无图片"), MB_OK)
		Config2Json(C.SleepTime, C.ChangeLockWallPaper)
	}
}

func setLockWallpaper(filepath string) error {
	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, LockWallPaperRegPath, registry.ALL_ACCESS)
	if err != nil {
		return errors.New("请使用管理员权限运行")
	}
	defer k.Close()

	err = k.SetStringValue("LockScreenImagePath", filepath)
	if err != nil {
		return errors.New("请使用管理员权限运行")
	}
	return nil
}

func UndoSetLockWallpaper() error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, LockWallPaperRegPath, registry.ALL_ACCESS)
	if err != nil {
		return errors.New("恢复锁屏失败,请使用管理员权限运行")
	}
	defer k.Close()

	err = k.DeleteValue("LockScreenImagePath")
	if err != nil {
		return errors.New("恢复锁屏失败,请使用管理员权限运行")
	}
	return nil
}

func EditConfig() {
	cmd := exec.Command("notepad", "./setting.json")
	cmd.Run()
}

func Config2Json(sleepTime int64, changeLock bool) {
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
	C.FolderPath = PicFolderPath
	SetTreeNode()
	cfg := Config{
		SleepTime:          sleepTime,
		ChangeLockWallPaper: changeLock,
		FolderPath:         PicFolderPath,
	}
	err := SaveData(cfg)
	if err != nil {
		return
	}
}

func SaveData(cfg Config) error {
	bytes, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile("./setting.json", bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadData() {
	file, err := os.ReadFile("setting.json")
	if err != nil {
		ShowMessage(errors.New("创建默认setting.json失败"), MB_OK)
		os.Exit(0)
		return
	}
	err = json.Unmarshal(file, &C)
	if err != nil {
		ShowMessage(errors.New("json文件解析失败"), MB_OK)
		os.Exit(1)
		return
	}
}
