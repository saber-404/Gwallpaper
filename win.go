/*
与win10图形功能相关的函数放在这里
*/

package Gwallpaper

import (
	"errors"
	"syscall"
	"unsafe"
)

// MessageBox constants
const (
	MB_OK                   = 0x00000000
	MB_OKCANCEL             = 0x00000001
	MB_ABORTRETRYIGNORE     = 0x00000002
	MB_YESNOCANCEL          = 0x00000003
	MB_YESNO                = 0x00000004
	MB_RETRYCANCEL          = 0x00000005
	MB_CANCELTRYCONTINUE    = 0x00000006
	MB_ICONHAND             = 0x00000010
	MB_ICONQUESTION         = 0x00000020
	MB_ICONEXCLAMATION      = 0x00000030
	MB_ICONASTERISK         = 0x00000040
	MB_USERICON             = 0x00000080
	MB_ICONWARNING          = MB_ICONEXCLAMATION
	MB_ICONERROR            = MB_ICONHAND
	MB_ICONINFORMATION      = MB_ICONASTERISK
	MB_ICONSTOP             = MB_ICONHAND
	MB_DEFBUTTON1           = 0x00000000
	MB_DEFBUTTON2           = 0x00000100
	MB_DEFBUTTON3           = 0x00000200
	MB_DEFBUTTON4           = 0x00000300
	MB_APPLMODAL            = 0x00000000
	MB_SYSTEMMODAL          = 0x00001000
	MB_TASKMODAL            = 0x00002000
	MB_HELP                 = 0x00004000
	MB_SETFOREGROUND        = 0x00010000
	MB_DEFAULT_DESKTOP_ONLY = 0x00020000
	MB_TOPMOST              = 0x00040000
	MB_RIGHT                = 0x00080000
	MB_RTLREADING           = 0x00100000
	MB_SERVICE_NOTIFICATION = 0x00200000
)

// BIF
const (
	BIF_RETURNONLYFSDIRS    = 0x00000001
	BIF_DONTGOBELOWDOMAIN   = 0x00000002
	BIF_STATUSTEXT          = 0x00000004
	BIF_RETURNFSANCESTORS   = 0x00000008
	BIF_EDITBOX             = 0x00000010
	BIF_VALIDATE            = 0x00000020
	BIF_NEWDIALOGSTYLE      = 0x00000040
	BIF_BROWSEINCLUDEURLS   = 0x00000080
	BIF_UAHINT              = 0x00000100
	BIF_NONEWFOLDERBUTTON   = 0x00000200
	BIF_NOTRANSLATETARGETS  = 0x00000400
	BIF_BROWSEFORCOMPUTER   = 0x00001000
	BIF_BROWSEFORPRINTER    = 0x00002000
	BIF_BROWSEINCLUDEFILES  = 0x00004000
	BIF_SHAREABLE           = 0x00008000
	BIF_BROWSEFILEJUNCTIONS = 0x00010000
)

// CSIDL
const (
	CSIDL_DESKTOP                 = 0x00
	CSIDL_INTERNET                = 0x01
	CSIDL_PROGRAMS                = 0x02
	CSIDL_CONTROLS                = 0x03
	CSIDL_PRINTERS                = 0x04
	CSIDL_PERSONAL                = 0x05
	CSIDL_FAVORITES               = 0x06
	CSIDL_STARTUP                 = 0x07
	CSIDL_RECENT                  = 0x08
	CSIDL_SENDTO                  = 0x09
	CSIDL_BITBUCKET               = 0x0A
	CSIDL_STARTMENU               = 0x0B
	CSIDL_MYDOCUMENTS             = 0x0C
	CSIDL_MYMUSIC                 = 0x0D
	CSIDL_MYVIDEO                 = 0x0E
	CSIDL_DESKTOPDIRECTORY        = 0x10
	CSIDL_DRIVES                  = 0x11
	CSIDL_NETWORK                 = 0x12
	CSIDL_NETHOOD                 = 0x13
	CSIDL_FONTS                   = 0x14
	CSIDL_TEMPLATES               = 0x15
	CSIDL_COMMON_STARTMENU        = 0x16
	CSIDL_COMMON_PROGRAMS         = 0x17
	CSIDL_COMMON_STARTUP          = 0x18
	CSIDL_COMMON_DESKTOPDIRECTORY = 0x19
	CSIDL_APPDATA                 = 0x1A
	CSIDL_PRINTHOOD               = 0x1B
	CSIDL_LOCAL_APPDATA           = 0x1C
	CSIDL_ALTSTARTUP              = 0x1D
	CSIDL_COMMON_ALTSTARTUP       = 0x1E
	CSIDL_COMMON_FAVORITES        = 0x1F
	CSIDL_INTERNET_CACHE          = 0x20
	CSIDL_COOKIES                 = 0x21
	CSIDL_HISTORY                 = 0x22
	CSIDL_COMMON_APPDATA          = 0x23
	CSIDL_WINDOWS                 = 0x24
	CSIDL_SYSTEM                  = 0x25
	CSIDL_PROGRAM_FILES           = 0x26
	CSIDL_MYPICTURES              = 0x27
	CSIDL_PROFILE                 = 0x28
	CSIDL_SYSTEMX86               = 0x29
	CSIDL_PROGRAM_FILESX86        = 0x2A
	CSIDL_PROGRAM_FILES_COMMON    = 0x2B
	CSIDL_PROGRAM_FILES_COMMONX86 = 0x2C
	CSIDL_COMMON_TEMPLATES        = 0x2D
	CSIDL_COMMON_DOCUMENTS        = 0x2E
	CSIDL_COMMON_ADMINTOOLS       = 0x2F
	CSIDL_ADMINTOOLS              = 0x30
	CSIDL_CONNECTIONS             = 0x31
	CSIDL_COMMON_MUSIC            = 0x35
	CSIDL_COMMON_PICTURES         = 0x36
	CSIDL_COMMON_VIDEO            = 0x37
	CSIDL_RESOURCES               = 0x38
	CSIDL_RESOURCES_LOCALIZED     = 0x39
	CSIDL_COMMON_OEM_LINKS        = 0x3A
	CSIDL_CDBURN_AREA             = 0x3B
	CSIDL_COMPUTERSNEARME         = 0x3D
	CSIDL_FLAG_CREATE             = 0x8000
	CSIDL_FLAG_DONT_VERIFY        = 0x4000
	CSIDL_FLAG_NO_ALIAS           = 0x1000
	CSIDL_FLAG_PER_USER_INIT      = 0x8000
	CSIDL_FLAG_MASK               = 0xFF00
)

// SPI
const (
	SPI_GETBEEP                     = 0x0001
	SPI_SETBEEP                     = 0x0002
	SPI_GETMOUSE                    = 0x0003
	SPI_SETMOUSE                    = 0x0004
	SPI_GETBORDER                   = 0x0005
	SPI_SETBORDER                   = 0x0006
	SPI_GETKEYBOARDSPEED            = 0x000a
	SPI_SETKEYBOARDSPEED            = 0x000b
	SPI_LANGDRIVER                  = 0x000c
	SPI_ICONHORIZONTALSPACING       = 0x000d
	SPI_GETSCREENSAVETIMEOUT        = 0x000e
	SPI_SETSCREENSAVETIMEOUT        = 0x000f
	SPI_GETSCREENSAVEACTIVE         = 0x0010
	SPI_SETSCREENSAVEACTIVE         = 0x0011
	SPI_GETGRIDGRANULARITY          = 0x0012
	SPI_SETGRIDGRANULARITY          = 0x0013
	SPI_SETDESKWALLPAPER            = 0x0014
	SPI_SETDESKPATTERN              = 0x0015
	SPI_GETKEYBOARDDELAY            = 0x0016
	SPI_SETKEYBOARDDELAY            = 0x0017
	SPI_ICONVERTICALSPACING         = 0x0018
	SPI_GETICONTITLEWRAP            = 0x0019
	SPI_SETICONTITLEWRAP            = 0x001a
	SPI_GETMENUDROPALIGNMENT        = 0x001b
	SPI_SETMENUDROPALIGNMENT        = 0x001c
	SPI_SETDOUBLECLKWIDTH           = 0x001d
	SPI_SETDOUBLECLKHEIGHT          = 0x001e
	SPI_GETICONTITLELOGFONT         = 0x001f
	SPI_SETDOUBLECLICKTIME          = 0x0020
	SPI_SETMOUSEBUTTONSWAP          = 0x0021
	SPI_SETICONTITLELOGFONT         = 0x0022
	SPI_GETFASTTASKSWITCH           = 0x0023
	SPI_SETFASTTASKSWITCH           = 0x0024
	SPI_SETDRAGFULLWINDOWS          = 0x0025
	SPI_GETDRAGFULLWINDOWS          = 0x0026
	SPI_GETNONCLIENTMETRICS         = 0x0029
	SPI_SETNONCLIENTMETRICS         = 0x002a
	SPI_GETMINIMIZEDMETRICS         = 0x002b
	SPI_SETMINIMIZEDMETRICS         = 0x002c
	SPI_GETICONMETRICS              = 0x002d
	SPI_SETICONMETRICS              = 0x002e
	SPI_SETWORKAREA                 = 0x002f
	SPI_GETWORKAREA                 = 0x0030
	SPI_SETPENWINDOWS               = 0x0031
	SPI_GETHIGHCONTRAST             = 0x0042
	SPI_SETHIGHCONTRAST             = 0x0043
	SPI_GETKEYBOARDPREF             = 0x0044
	SPI_SETKEYBOARDPREF             = 0x0045
	SPI_GETSCREENREADER             = 0x0046
	SPI_SETSCREENREADER             = 0x0047
	SPI_GETANIMATION                = 0x0048
	SPI_SETANIMATION                = 0x0049
	SPI_GETFONTSMOOTHING            = 0x004a
	SPI_SETFONTSMOOTHING            = 0x004b
	SPI_SETDRAGWIDTH                = 0x004c
	SPI_SETDRAGHEIGHT               = 0x004d
	SPI_SETHANDHELD                 = 0x004e
	SPI_GETLOWPOWERTIMEOUT          = 0x004f
	SPI_GETPOWEROFFTIMEOUT          = 0x0050
	SPI_SETLOWPOWERTIMEOUT          = 0x0051
	SPI_SETPOWEROFFTIMEOUT          = 0x0052
	SPI_GETLOWPOWERACTIVE           = 0x0053
	SPI_GETPOWEROFFACTIVE           = 0x0054
	SPI_SETLOWPOWERACTIVE           = 0x0055
	SPI_SETPOWEROFFACTIVE           = 0x0056
	SPI_SETCURSORS                  = 0x0057
	SPI_SETICONS                    = 0x0058
	SPI_GETDEFAULTINPUTLANG         = 0x0059
	SPI_SETDEFAULTINPUTLANG         = 0x005a
	SPI_SETLANGTOGGLE               = 0x005b
	SPI_GETWINDOWSEXTENSION         = 0x005c
	SPI_SETMOUSETRAILS              = 0x005d
	SPI_GETMOUSETRAILS              = 0x005e
	SPI_SETSCREENSAVERRUNNING       = 0x0061
	SPI_SCREENSAVERRUNNING          = SPI_SETSCREENSAVERRUNNING
	SPI_GETFILTERKEYS               = 0x0032
	SPI_SETFILTERKEYS               = 0x0033
	SPI_GETTOGGLEKEYS               = 0x0034
	SPI_SETTOGGLEKEYS               = 0x0035
	SPI_GETMOUSEKEYS                = 0x0036
	SPI_SETMOUSEKEYS                = 0x0037
	SPI_GETSHOWSOUNDS               = 0x0038
	SPI_SETSHOWSOUNDS               = 0x0039
	SPI_GETSTICKYKEYS               = 0x003a
	SPI_SETSTICKYKEYS               = 0x003b
	SPI_GETACCESSTIMEOUT            = 0x003c
	SPI_SETACCESSTIMEOUT            = 0x003d
	SPI_GETSERIALKEYS               = 0x003e
	SPI_SETSERIALKEYS               = 0x003f
	SPI_GETSOUNDSENTRY              = 0x0040
	SPI_SETSOUNDSENTRY              = 0x0041
	SPI_GETSNAPTODEFBUTTON          = 0x005f
	SPI_SETSNAPTODEFBUTTON          = 0x0060
	SPI_GETMOUSEHOVERWIDTH          = 0x0062
	SPI_SETMOUSEHOVERWIDTH          = 0x0063
	SPI_GETMOUSEHOVERHEIGHT         = 0x0064
	SPI_SETMOUSEHOVERHEIGHT         = 0x0065
	SPI_GETMOUSEHOVERTIME           = 0x0066
	SPI_SETMOUSEHOVERTIME           = 0x0067
	SPI_GETWHEELSCROLLLINES         = 0x0068
	SPI_SETWHEELSCROLLLINES         = 0x0069
	SPI_GETMENUSHOWDELAY            = 0x006a
	SPI_SETMENUSHOWDELAY            = 0x006b
	SPI_GETWHEELSCROLLCHARS         = 0x006c
	SPI_SETWHEELSCROLLCHARS         = 0x006d
	SPI_GETSHOWIMEUI                = 0x006e
	SPI_SETSHOWIMEUI                = 0x006f
	SPI_GETMOUSESPEED               = 0x0070
	SPI_SETMOUSESPEED               = 0x0071
	SPI_GETSCREENSAVERRUNNING       = 0x0072
	SPI_GETDESKWALLPAPER            = 0x0073
	SPI_GETAUDIODESCRIPTION         = 0x0074
	SPI_SETAUDIODESCRIPTION         = 0x0075
	SPI_GETSCREENSAVESECURE         = 0x0076
	SPI_SETSCREENSAVESECURE         = 0x0077
	SPI_GETHUNGAPPTIMEOUT           = 0x0078
	SPI_SETHUNGAPPTIMEOUT           = 0x0079
	SPI_GETWAITTOKILLTIMEOUT        = 0x007a
	SPI_SETWAITTOKILLTIMEOUT        = 0x007b
	SPI_GETWAITTOKILLSERVICETIMEOUT = 0x007c
	SPI_SETWAITTOKILLSERVICETIMEOUT = 0x007d
	SPI_GETMOUSEDOCKTHRESHOLD       = 0x007e
	SPI_SETMOUSEDOCKTHRESHOLD       = 0x007f
	SPI_GETPENDOCKTHRESHOLD         = 0x0080
	SPI_SETPENDOCKTHRESHOLD         = 0x0081
	SPI_GETWINARRANGING             = 0x0082
	SPI_SETWINARRANGING             = 0x0083
	SPI_GETMOUSEDRAGOUTTHRESHOLD    = 0x0084
	SPI_SETMOUSEDRAGOUTTHRESHOLD    = 0x0085
	SPI_GETPENDRAGOUTTHRESHOLD      = 0x0086
	SPI_SETPENDRAGOUTTHRESHOLD      = 0x0087
	SPI_GETMOUSESIDEMOVETHRESHOLD   = 0x0088
	SPI_SETMOUSESIDEMOVETHRESHOLD   = 0x0089
	SPI_GETPENSIDEMOVETHRESHOLD     = 0x008a
	SPI_SETPENSIDEMOVETHRESHOLD     = 0x008b
	SPI_GETDRAGFROMMAXIMIZE         = 0x008c
	SPI_SETDRAGFROMMAXIMIZE         = 0x008d
	SPI_GETSNAPSIZING               = 0x008e
	SPI_SETSNAPSIZING               = 0x008f
	SPI_GETDOCKMOVING               = 0x0090
	SPI_SETDOCKMOVING               = 0x0091
)

type (
	HWND    uintptr
	CSIDL   uint32
	BOOL    int32
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

	shell32                = syscall.NewLazyDLL("shell32.dll")
	shBrowseForFolder      = shell32.NewProc("SHBrowseForFolderW")
	shGetPathFromIDList    = shell32.NewProc("SHGetPathFromIDListW")
	shGetSpecialFolderPath = shell32.NewProc("SHGetSpecialFolderPathW")
	shParseDisplayName     = shell32.NewProc("SHParseDisplayName")
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

// SHGetSpecialFolderPath 解析特殊文件夹，如桌面
func SHGetSpecialFolderPath(hwndOwner HWND, lpszPath *uint16, csidl CSIDL, fCreate bool) bool {
	ret, _, _ := syscall.Syscall6(shGetSpecialFolderPath.Addr(), 4,
		uintptr(hwndOwner),
		uintptr(unsafe.Pointer(lpszPath)),
		uintptr(csidl),
		uintptr(BoolToBOOL(fCreate)),
		0,
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
		//PszDisplayName: syscall.StringToUTF16Ptr(Title),
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
