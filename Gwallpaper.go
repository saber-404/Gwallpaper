package Gwallpaper

import (
	_ "embed"
	"encoding/json"
	"errors"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	C        Config
	Prefix   string
	TreeNode PicNode

	//go:embed asset/icon.ico
	Icon []byte
)

const (
	Title                = "GwallPaper"
	LockWallPaperRegPath = `SOFTWARE\Microsoft\Windows\CurrentVersion\PersonalizationCSP`
)

type Config struct {
	FolderPath         string `json:"FolderPath"`
	SleepTime          int64  `json:"SleepTime"`
	ChangLockWallPaper bool   `json:"ChangLockWallPaper"`
}

func init() {
	InitSetting()
}

// ChangeWallPaper 改变壁纸
func (c *Config) ChangeWallPaper() {
	//path := C.GetPicPath()
	path := C.GetPicPathByTree()
	if !IsImage(path) {
		//	更新缓存
		//c.SetPrefixAndPicPath()
		SetTreeNode()
	}
	err := SetWallpaper(path)
	if err != nil {
		ShowMessage(err, MB_OK)
		return
	}
	if c.ChangLockWallPaper {
		err := setLockWallpaper(path)
		if err != nil {
			ShowMessage(err, MB_OK)
			return
		}
	}
	//	测试
	//logt := fmt.Sprintf("Prefix:%s Pics:%v", Prefix, PicPath)
	//ShowMessage(errors.New(logt), MB_OK)

	//	测试
	//PrintTree(&TreeNode, 0)
}

// InitSetting 加载配置
func InitSetting() {
	_, err := os.Stat("setting.json")
	if err != nil {
		Config2Json()
	}
	file, err := ioutil.ReadFile("setting.json")
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
	//扫描一次图片,并缓存到变量
	//C.SetPrefixAndPicPath()
	SetTreeNode()
}

// 锁屏壁纸设置
func setLockWallpaper(filepath string) error {
	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, LockWallPaperRegPath, registry.ALL_ACCESS)
	if err != nil {
		return errors.New("请使用管理员权限运行")
	}
	defer k.Close()

	// Set the value of LockScreenImagePath to the desired path
	err = k.SetStringValue("LockScreenImagePath", filepath)
	if err != nil {
		return errors.New("请使用管理员权限运行")
	}
	return nil
}

// UndoSetLockWallpaper 撤销锁屏壁纸设置
func UndoSetLockWallpaper() error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, LockWallPaperRegPath, registry.ALL_ACCESS)
	if err != nil {
		return errors.New("恢复锁屏失败,请使用管理员权限运行")
	}
	defer k.Close()

	// Delete the value of LockScreenImagePath
	err = k.DeleteValue("LockScreenImagePath")
	if err != nil {
		return errors.New("恢复锁屏失败,请使用管理员权限运行")
	}
	return nil
}

// EditConfig 编辑配置
func EditConfig() {
	cmd := exec.Command("notepad", "./setting.json")
	cmd.Run()
}

// Config2Json 生成配置 也可用于还原原本配置
func Config2Json() {
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
	DefaultConfig := Config{
		FolderPath:         PicFolderPath + `\`,
		SleepTime:          900,
		ChangLockWallPaper: false,
	}
	bytes, err := json.Marshal(DefaultConfig)
	if err != nil {
		return
	}
	err = ioutil.WriteFile("./setting.json", bytes, 0644)
	if err != nil {
		return
	}
}
