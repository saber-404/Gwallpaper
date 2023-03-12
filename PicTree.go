package Gwallpaper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type PicNode struct {
	name     string
	children []*PicNode
}

// Insert 生成图片树
func (p *PicNode) Insert(path string) {
	// 递归出口
	dirs, err := os.ReadDir(path)
	if err != nil {
		return
	}
	dirs = Filter(path, dirs)
	for _, file := range dirs {
		node := &PicNode{name: file.Name()}
		p.children = append(p.children, node)
		node.Insert(filepath.Join(path, node.name))
	}

}

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
	TreeNode.children = nil
	TreeNode.Insert(C.FolderPath)
}

// GetPicPathByTree	随机返回一张图片的绝对路径
func (c *Config) GetPicPathByTree() string {
	var path string
	node := &TreeNode
	for {
		l := len(node.children)
		if l == 0 {
			break
		}
		child := node.children[RandIntn(l)]
		path = filepath.Join(path, child.name)
		node = child
	}
	return filepath.Join(C.FolderPath, path)
}

func PrintTree(node *PicNode, depth int) {
	// 打印节点的名称和缩进
	fmt.Printf("%s%s\n", strings.Repeat(" ", depth*4), node.name)
	// 递归打印子节点
	for _, child := range node.children {
		PrintTree(child, depth+1)
	}
}
