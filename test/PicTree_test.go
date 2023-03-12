package test

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type PicNode struct {
	name     string
	children []*PicNode
}

type PicTree struct {
	rootPath string
	root     *PicNode
}

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

func printNode(node *PicNode, depth int) {
	// 打印节点的名称和缩进
	fmt.Printf("%s%s\n", strings.Repeat(" ", depth*4), node.name)
	// 递归打印子节点
	for _, child := range node.children {
		printNode(child, depth+1)
	}
}

func TestF(t *testing.T) {
	tree := PicTree{}
	tree.rootPath = `D:\datacenter\壁纸\ForWallPaper`
	root := &PicNode{}
	tree.root = root
	tree.root.Insert(tree.rootPath)
	// 测试多次调用 printTree 函数
	for i := 0; i < 10; i++ {
		//rand.Seed(time.Now().UnixNano())
		fmt.Println(printTree(tree))
	}
}

func Filter(path string, dirs []os.DirEntry) []os.DirEntry {
	var result []os.DirEntry
	for _, dir := range dirs {
		if dir.IsDir() {
			empty, err := IsDirEmpty(filepath.Join(path, dir.Name()))
			if err != nil {
				continue
			}
			if empty {
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

func IsImage(path string) bool {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return false
	}
	_, _, err = image.DecodeConfig(file)
	if err != nil {
		return false
	}
	return true
}

func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func printTree(tree PicTree) string {
	//	随机返回一张图片的绝对路径
	//rand.Seed(time.Now().UnixNano())
	var path string
	node := tree.root
	for {
		if len(node.children) == 0 {
			break
		}
		child := node.children[RandIntn(len(node.children))]
		path = filepath.Join(path, child.name)
		node = child
	}
	return filepath.Join(tree.rootPath, path)
}
func RandIntn(length int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(length)
}
