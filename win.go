/*
与win10图形功能相关的函数放在这里
*/

package Gwallpaper

import (
	"errors"
	"syscall"
	"unsafe"
)

const MB_OK = 0x00000000

const (
	BIF_RETURNONLYFSDIRS  = 0x00000001
	BIF_DONTGOBELOWDOMAIN  = 0x00000002
	BIF_NEWDIALOGSTYLE     = 0x00000040
	BIF_NONEWFOLDERBUTTON  = 0x00000200
)

const SPI_SETDESKWALLPAPER = 0x0014

type (
	HWND    uintptr
	HRESULT int32
)

type SHITEMID struct {
	CB   uint16
	ABID [1]byte
}

type ITEMIDLIST struct {
	ID SHITEMID
}

type BROWSEINFO struct {
	HwndOwner      HWND
	PidlRoot       uintptr
	PszDisplayName *uint16
	LpszTitle      *uint16
	UlFlags        uint32
	Lpfn           uintptr
	LParam         uintptr
	IImage         int32
}

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	systemParametersInfo = user32.NewProc("SystemParametersInfoW")
	messageBox           = user32.NewProc("MessageBoxW")
	getActiveWindow      = user32.NewProc("GetActiveWindow")

	shell32             = syscall.NewLazyDLL("shell32.dll")
	shBrowseForFolder   = shell32.NewProc("SHBrowseForFolderW")
	shGetPathFromIDList = shell32.NewProc("SHGetPathFromIDListW")
	shParseDisplayName  = shell32.NewProc("SHParseDisplayName")
)

// GetActiveWindow 获得一个句柄
func GetActiveWindow() HWND {
	ret, _, _ := syscall.Syscall(getActiveWindow.Addr(), 0,
		0,
		0,
		0)

	return HWND(ret)
}

// MessageBox 弹出对话框
func MessageBox(hWnd HWND, lpText, lpCaption *uint16, uType uint32) int32 {
	ret, _, _ := syscall.Syscall6(messageBox.Addr(), 4,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		uintptr(uType),
		0,
		0)

	return int32(ret)
}

// ShowMessage 显示对话框
func ShowMessage(err error, flags uintptr) {
	hwnd := GetActiveWindow()
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
	MessageBox(hwnd, messageptr, captionptr, uint32(flags))
}

// SHBrowseForFolder 返回文件夹选择器对象指针
func SHBrowseForFolder(lpbi *BROWSEINFO) uintptr {
	ret, _, _ := syscall.Syscall(shBrowseForFolder.Addr(), 1,
		uintptr(unsafe.Pointer(lpbi)),
		0,
		0)
	return ret
}

// SHGetPathFromIDList 返回文件夹路径
func SHGetPathFromIDList(pidl uintptr, pszPath *uint16) bool {
	ret, _, _ := syscall.Syscall(shGetPathFromIDList.Addr(), 2,
		pidl,
		uintptr(unsafe.Pointer(pszPath)),
		0)

	return ret != 0
}

// SHParseDisplayName 解析路径到结构体
func SHParseDisplayName(pszName *uint16, pbc uintptr, ppidl *uintptr, sfgaoIn uint32, psfgaoOut *uint32) HRESULT {
	ret, _, _ := syscall.Syscall6(shParseDisplayName.Addr(), 5,
		uintptr(unsafe.Pointer(pszName)),
		pbc,
		uintptr(unsafe.Pointer(ppidl)),
		0,
		uintptr(unsafe.Pointer(psfgaoOut)),
		0)

	return HRESULT(ret)
}

// ShowFolderDialogForGetFolderPath 显示文件夹选择器并返回地址
func ShowFolderDialogForGetFolderPath(message string) (IsChoice bool, PicFolderPath string) {
	var path [256]uint16
	//选择我的电脑为起始文件夹
	pszName := syscall.StringToUTF16Ptr(`::{20D04FE0-3AEA-1069-A2D8-08002B30309D}`)
	ppidl := uintptr(unsafe.Pointer(&ITEMIDLIST{}))
	SHParseDisplayName(pszName, 0, &ppidl, 0, nil)
	bi := &BROWSEINFO{
		HwndOwner: HWND(0),
		PidlRoot:  ppidl,
		LpszTitle: syscall.StringToUTF16Ptr(message),
		UlFlags:   BIF_RETURNONLYFSDIRS | BIF_DONTGOBELOWDOMAIN | BIF_NEWDIALOGSTYLE | BIF_NONEWFOLDERBUTTON,
	}
	pidl := SHBrowseForFolder(bi)
	IsChoice = SHGetPathFromIDList(pidl, &path[0])
	return IsChoice, syscall.UTF16ToString(path[:])
}

func SystemParametersInfo(uiAction, uiParam uint32, pvParam unsafe.Pointer, fWinIni uint32) bool {
	ret, _, _ := syscall.Syscall6(systemParametersInfo.Addr(), 4,
		uintptr(uiAction),
		uintptr(uiParam),
		uintptr(pvParam),
		uintptr(fWinIni),
		0,
		0)
	return ret != 0
}

// SetWallpaper 桌面壁纸设置函数
func SetWallpaper(filepath string) error {
	// 将文件路径转换为指向宽字符的指针
	filepathPtr, err := syscall.UTF16PtrFromString(filepath)
	if err != nil {
		return errors.New("文件路径转换为指向宽字符的指针失败")
	}
	IsSet := SystemParametersInfo(
		SPI_SETDESKWALLPAPER,
		0,
		unsafe.Pointer(filepathPtr),
		2,
	)
	if !IsSet {
		return errors.New("设置壁纸失败")
	}
	return nil
}
