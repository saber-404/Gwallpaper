package wallpaper

import (
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
