package wallpaper

import (
	_ "embed"
	"errors"
	"golang.org/x/sys/windows/registry"
	"os/exec"
)

var (
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
