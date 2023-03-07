package test

import (
	"Gwallpaper"
	"errors"
	"fmt"
	"syscall"
	"testing"
	"unsafe"
)

// OFN
const (
	OFN_ALLOWMULTISELECT     = 0x00000200
	OFN_CREATEPROMPT         = 0x00002000
	OFN_DONTADDTORECENT      = 0x02000000
	OFN_ENABLEHOOK           = 0x00000020
	OFN_ENABLEINCLUDENOTIFY  = 0x00400000
	OFN_ENABLESIZING         = 0x00800000
	OFN_ENABLETEMPLATE       = 0x00000040
	OFN_ENABLETEMPLATEHANDLE = 0x00000080
	OFN_EXPLORER             = 0x00080000
	OFN_EXTENSIONDIFFERENT   = 0x00000400
	OFN_FILEMUSTEXIST        = 0x00001000
	OFN_FORCESHOWHIDDEN      = 0x10000000
	OFN_HIDEREADONLY         = 0x00000004
	OFN_LONGNAMES            = 0x00200000
	OFN_NOCHANGEDIR          = 0x00000008
	OFN_NODEREFERENCELINKS   = 0x00100000
	OFN_NOLONGNAMES          = 0x00040000
	OFN_NONETWORKBUTTON      = 0x00020000
	OFN_NOREADONLYRETURN     = 0x00008000
	OFN_NOTESTFILECREATE     = 0x00010000
	OFN_NOVALIDATE           = 0x00000100
	OFN_OVERWRITEPROMPT      = 0x00000002
	OFN_PATHMUSTEXIST        = 0x00000800
	OFN_READONLY             = 0x00000001
	OFN_SHAREAWARE           = 0x00004000
	OFN_SHOWHELP             = 0x00000010
)

var (
	comdlg32        = syscall.NewLazyDLL("comdlg32.dll")
	getOpenFileName = comdlg32.NewProc("GetOpenFileNameW")
)

type OPENFILENAME struct {
	LStructSize       uint32
	HwndOwner         uintptr
	HInstance         uintptr
	LpstrFilter       *uint16
	LpstrCustomFilter *uint16
	NMaxCustFilter    uint32
	NFilterIndex      uint32
	LpstrFile         *uint16
	NMaxFile          uint32
	LpstrFileTitle    *uint16
	NMaxFileTitle     uint32
	LpstrInitialDir   *uint16
	LpstrTitle        *uint16
	Flags             uint32
	NFileOffset       uint16
	NFileExtension    uint16
	LpstrDefExt       *uint16
	LCustData         uintptr
	LpfnHook          uintptr
	LpTemplateName    *uint16
	PvReserved        unsafe.Pointer
	DwReserved        uint32
	FlagsEx           uint32
}

// 通过文件选择对话框获取文件地址
func MyOpenFileDialog() {
	var folder [256]uint16
	ofn := &OPENFILENAME{
		LStructSize: uint32(unsafe.Sizeof(OPENFILENAME{})),
		HwndOwner:   uintptr(0),
		LpstrFilter: syscall.StringToUTF16Ptr(""),
		LpstrFile:   &folder[0],
		NMaxFile:    uint32(len(folder)),
		//Flags:       OFN_FILEMUSTEXIST | OFN_PATHMUSTEXIST,
		Flags: OFN_FILEMUSTEXIST | OFN_PATHMUSTEXIST,
	}
	ret, _, err := getOpenFileName.Call(uintptr(unsafe.Pointer(ofn)))
	if ret == 0 {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Selected folder: %v\n", syscall.UTF16ToString(folder[:]))
}

func TestMyMyOpenFileDialog(t *testing.T) {
	MyOpenFileDialog()
}

func TestShowMessage(t *testing.T) {
	err_test := errors.New("test error")
	Gwallpaper.ShowMessage(err_test, Gwallpaper.MB_OK)
}

func TestShowFolderDialogForGetFolderPath(t *testing.T) {
	IsChoice, picpath := Gwallpaper.ShowFolderDialogForGetFolderPath("选择壁纸文件夹")
	if IsChoice {
		fmt.Printf("用户做出选择了，文件夹路径是 %s\n", picpath)
	} else {
		fmt.Printf("用户没有做出选择")
	}
}

func TestSetWallpaper(t *testing.T) {
	newWallpaperPath := `D:\datacenter\壁纸\ForWallPaper\4Desktop\105978873_p0.png`
	err := Gwallpaper.SetWallpaper(newWallpaperPath)
	if err != nil {
		t.Log(err)
		return
	}
}
