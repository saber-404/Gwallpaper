package test

import (
	"Gwallpaper"
	"testing"
)

func TestReload(t *testing.T) {
	//Gwallpaper.InitSetting()
	Gwallpaper.LoadData()
	Gwallpaper.SetTreeNode()
	Gwallpaper.C.ChangeWallPaper()
	Gwallpaper.SaveData(Gwallpaper.C)
}

/*var (
	C        Config
	TreeNode PicNode
)

type PicNode struct {
	Name     string
	Children []*PicNode
}
type Config struct {
	SleepTime          int64
	ChangLockWallPaper bool
	Cache              *PicNode
}

func LoadData() {
	file, err := ioutil.ReadFile("setting.json")
	if err != nil {
		//ShowMessage(errors.New("创建默认setting.json失败"), MB_OK)
		//os.Exit(0)
		return
	}
	err = json.Unmarshal(file, &C)
	if err != nil {
		//ShowMessage(errors.New("json文件解析失败"), MB_OK)
		//os.Exit(1)
		return
	}
}*/
