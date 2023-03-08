package Gwallpaper

import (
	"encoding/json"
	"errors"
	"golang.org/x/sys/windows/registry"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	systemParametersInfo = user32.NewProc("SystemParametersInfoW")
	spiSetDeskWallpaper  = uintptr(20)
	messageBox           = user32.NewProc("MessageBoxW")
	getActiveWindow      = user32.NewProc("GetActiveWindow")
	C                    Config
)

const (
	Title                = "GwallPaper"
	LockWallPaperRegPath = `SOFTWARE\Microsoft\Windows\CurrentVersion\PersonalizationCSP`
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
)

type Config struct {
	RetryTimes         int    `json:"RetryTimes"`
	FolderPath         string `json:"FolderPath"`
	SleepTime          int64  `json:"SleepTime"`
	ChangLockWallPaper bool   `json:"ChangLockWallPaper"`
}

func init() {
	InitSetting()
}

// ChangeWallPaper 改变壁纸
func (c *Config) ChangeWallPaper() {
	PicName, err := C.GetPicName()
	if err != nil {
		ShowMessage(err, MB_OK)
		return
	}
	err = setWallpaper(C.FolderPath + PicName)
	if err != nil {
		ShowMessage(err, MB_OK)
		return
	}
	if c.ChangLockWallPaper {
		err := setLockWallpaper(C.FolderPath + PicName)
		if err != nil {
			ShowMessage(err, MB_OK)
			return
		}
	}
}

// GetPicName 获取图片名称
func (c *Config) GetPicName() (PicName string, err error) {
	files, err := os.ReadDir(c.FolderPath)
	if err != nil {
		return
	}
	rand.Seed(time.Now().UnixNano())
	retry := 0
	index := rand.Intn(len(files))
	for !isImage(files[index]) {
		rand.Seed(time.Now().UnixNano())
		index = rand.Intn(len(files))
		retry += 1
		if retry == c.RetryTimes {
			return "", errors.New("the folder not has pic")
		}
	}
	return files[index].Name(), nil
}

// 判断是否是图片
func isImage(file fs.DirEntry) bool {
	if file.IsDir() {
		return false
	}
	ext := strings.ToLower(filepath.Ext(file.Name()))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

// GetIcon icon转byte流
func GetIcon(path string) (iconbytes []byte) {
	iconbytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return iconbytes
}

// ShowMessage 显示对话框
func ShowMessage(err error, flags uintptr) {
	var getActiveWindowFunc = getActiveWindow.Addr()
	var hwnd uintptr
	ret, _, _ := syscall.SyscallN(getActiveWindowFunc, 0, 0, 0, 0)
	if ret != 0 {
		hwnd = ret
	}

	var caption = Title
	var message = err.Error()

	messageptr, err := syscall.UTF16PtrFromString(message)
	if err != nil {
		return
	}
	captionptr, err := syscall.UTF16PtrFromString(caption)
	if err != nil {
		return
	}
	messageBox.Call(hwnd,
		uintptr(unsafe.Pointer(messageptr)),
		uintptr(unsafe.Pointer(captionptr)),
		flags)
}

// InitSetting 加载配置
func InitSetting() {
	file, err := ioutil.ReadFile("setting.json")
	if err != nil {
		ShowMessage(errors.New("找不到setting.json"), MB_OK)
		return
	}
	err = json.Unmarshal(file, &C)
	if err != nil {
		ShowMessage(errors.New("json文件解析失败"), MB_OK)
		return
	}
}

// SetWallpaper 桌面壁纸设置函数
func setWallpaper(filepath string) error {
	// 将文件路径转换为指向宽字符的指针
	filepathPtr, err := syscall.UTF16PtrFromString(filepath)
	if err != nil {
		return errors.New("文件路径转换为指向宽字符的指针失败")
	}
	// 调用 SystemParametersInfo 函数设置壁纸
	_, _, err = systemParametersInfo.Call(
		spiSetDeskWallpaper,
		0,
		uintptr(unsafe.Pointer(filepathPtr)),
		uintptr(2),
	)
	return nil
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

func EditConfig() {
	cmd := exec.Command("notepad", "./setting.json")
	cmd.Run()
}
