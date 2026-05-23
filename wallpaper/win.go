package wallpaper

import (
	"errors"
	"syscall"
	"unsafe"
)

const MB_OK = 0x00000000

const (
	BIF_RETURNONLYFSDIRS  = 0x00000001
	BIF_DONTGOBELOWDOMAIN = 0x00000002
	BIF_NEWDIALOGSTYLE    = 0x00000040
	BIF_NONEWFOLDERBUTTON = 0x00000200
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

func GetActiveWindow() HWND {
	ret, _, _ := syscall.SyscallN(getActiveWindow.Addr())
	return HWND(ret)
}

func MessageBox(hWnd HWND, lpText, lpCaption *uint16, uType uint32) int32 {
	ret, _, _ := syscall.SyscallN(messageBox.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		uintptr(uType))
	return int32(ret)
}

func ShowMessage(err error, flags uint32) {
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
	MessageBox(hwnd, messageptr, captionptr, flags)
}

func SHBrowseForFolder(lpbi *BROWSEINFO) uintptr {
	ret, _, _ := syscall.SyscallN(shBrowseForFolder.Addr(),
		uintptr(unsafe.Pointer(lpbi)))
	return ret
}

func SHGetPathFromIDList(pidl uintptr, pszPath *uint16) bool {
	ret, _, _ := syscall.SyscallN(shGetPathFromIDList.Addr(),
		pidl,
		uintptr(unsafe.Pointer(pszPath)))
	return ret != 0
}

func SHParseDisplayName(pszName *uint16, pbc uintptr, ppidl *uintptr, psfgaoOut *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(shParseDisplayName.Addr(),
		uintptr(unsafe.Pointer(pszName)),
		pbc,
		uintptr(unsafe.Pointer(ppidl)),
		uintptr(unsafe.Pointer(psfgaoOut)))
	return HRESULT(ret)
}

func ShowFolderDialogForGetFolderPath(message string) (IsChoice bool, PicFolderPath string) {
	var path [256]uint16
	pszName, err := syscall.UTF16PtrFromString(`::{20D04FE0-3AEA-1069-A2D8-08002B30309D}`)
	if err != nil {
		return false, ""
	}
	ppidl := uintptr(unsafe.Pointer(&ITEMIDLIST{}))
	SHParseDisplayName(pszName, 0, &ppidl, nil)
	titlePtr, _ := syscall.UTF16PtrFromString(message)
	bi := &BROWSEINFO{
		HwndOwner: HWND(0),
		PidlRoot:  ppidl,
		LpszTitle: titlePtr,
		UlFlags:   BIF_RETURNONLYFSDIRS | BIF_DONTGOBELOWDOMAIN | BIF_NEWDIALOGSTYLE | BIF_NONEWFOLDERBUTTON,
	}
	pidl := SHBrowseForFolder(bi)
	IsChoice = SHGetPathFromIDList(pidl, &path[0])
	return IsChoice, syscall.UTF16ToString(path[:])
}

func SystemParametersInfo(uiAction, uiParam uint32, pvParam unsafe.Pointer, fWinIni uint32) bool {
	ret, _, _ := syscall.SyscallN(systemParametersInfo.Addr(),
		uintptr(uiAction),
		uintptr(uiParam),
		uintptr(pvParam),
		uintptr(fWinIni))
	return ret != 0
}

func SetWallpaper(filepath string) error {
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
