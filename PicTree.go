package Gwallpaper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type PicNode struct {
	Name     string
	Children []*PicNode
}

// Insert 生成图片树
func (p *PicNode) Insert(path string) {
	// 递归出口
	dirs, err := os.ReadDir(path)
	if err != nil {
		return
	}
	dirs = Filter(path, dirs)
	if len(dirs) == 0 {
		return
	}
	for _, file := range dirs {
		node := &PicNode{Name: file.Name()}
		p.Children = append(p.Children, node)
		node.Insert(filepath.Join(path, node.Name))
	}

}

/*func (p *PicNode) SaveData2File(filePath string) {
	//	将p中的数据序列化到文件filePath
	data, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile(filePath, data, 0644)
}*/

/*func (p *PicNode) LoadDataFromFile(filePath string) error {
	//	从filePath中反序列化数据，重建树
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, p)
	if err != nil {
		return err
	}
	return nil
}*/

// Filter 去掉dirs中的空文件夹
func Filter(path string, dirs []os.DirEntry) []os.DirEntry {
	var result []os.DirEntry
	for _, dir := range dirs {
		if dir.IsDir() {
			if !CheckFolderHasImage(filepath.Join(path, dir.Name())) {
				continue
			}
		} else {
			if !IsImage(filepath.Join(path, dir.Name())) {
				continue
			}
		}
		result = append(result, dir)
	}
	return result
}

func SetTreeNode() {
	C.Cache.Children = nil
	//TreeNode.Children = nil
	C.Cache.Insert(C.Cache.Name)
}

// GetPicPathByTree	随机返回一张图片的绝对路径
func (c *Config) GetPicPathByTree() string {
	if len(C.Cache.Children) == 0 {
		ShowMessage(errors.New("壁纸文件夹内无图片"), MB_OK)
		Config2Json(C.SleepTime, C.ChangLockWallPaper)
	}
	var path string
	node := C.Cache
	for {
		l := len(node.Children)
		if l == 0 {
			break
		}
		child := node.Children[RandIntn(l)]
		path = filepath.Join(path, child.Name)
		node = *child
	}
	return filepath.Join(C.Cache.Name, path)
}

func PrintTree(node *PicNode, depth int) {
	// 打印节点的名称和缩进
	fmt.Printf("%s%s\n", strings.Repeat(" ", depth*4), node.Name)
	// 递归打印子节点
	for _, child := range node.Children {
		PrintTree(child, depth+1)
	}
}
