package wallpaper

import (
	"errors"
	"os"
	"path/filepath"
)

type PicNode struct {
	Name     string
	Children []*PicNode
}

func (p *PicNode) Insert(path string) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return
	}
	dirs = filterPicEntries(path, dirs)
	if len(dirs) == 0 {
		return
	}
	for _, file := range dirs {
		node := &PicNode{Name: file.Name()}
		p.Children = append(p.Children, node)
		node.Insert(filepath.Join(path, node.Name))
	}
}

func filterPicEntries(path string, dirs []os.DirEntry) []os.DirEntry {
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
	picTree.Children = nil
	picTree.Insert(C.FolderPath)
}

func GetPicPathByTree() string {
	if len(picTree.Children) == 0 {
		ShowMessage(errors.New("壁纸文件夹内无图片"), MB_OK)
		Config2Json(C.SleepTime, C.ChangeLockWallPaper)
	}
	var path string
	node := picTree
	for {
		l := len(node.Children)
		if l == 0 {
			break
		}
		child := node.Children[RandIntn(l)]
		path = filepath.Join(path, child.Name)
		node = *child
	}
	return filepath.Join(C.FolderPath, path)
}
