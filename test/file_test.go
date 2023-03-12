package test

import (
	"fmt"
	"testing"
)

// Node represents a node in the tree
type Node struct {
	value    rune    // character value
	children []*Node // slice of pointers to children nodes
}

// Tree represents a tree of strings
type Tree struct {
	root *Node // root node of the tree
}

// Insert inserts a string into the tree by creating nodes for each character if they do not exist
func (t *Tree) Insert(s string) {
	// if the root node is nil, create one with an empty value
	if t.root == nil {
		t.root = &Node{}
	}
	// start from the root node
	current := t.root
	// loop through each character in the string
	for _, c := range s {
		// flag to indicate if the character already exists as a child node
		found := false
		// loop through each child node of current node
		for _, child := range current.children {
			// if child node has same value as character, set current node to child node and break inner loop
			if child.value == c {
				current = child
				found = true
				break
			}
		}
		// if character does not exist as a child node, create one with character value and append it to current node's children slice
		if !found {
			newNode := &Node{value: c}
			current.children = append(current.children, newNode)
			current = newNode // set current node to new node
		}
	}
}

// Print prints a node's value and its children recursively using indentation
func Print(node *Node, level int) {
	// print indentation according to level
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	// print node's value (if not empty) followed by newline
	if node.value != 0 { // 0 is rune's zero value which means empty
		fmt.Println(string(node.value))
	} else {
		fmt.Println()
	}

	// loop through each child node and print it recursively with increased level
	for _, child := range node.children {
		Print(child, level+1)
	}
}

func TestIn(t1 *testing.T) {
	// initialize an empty tree
	t := &Tree{}
	// slice of strings to be inserted into the tree
	s := []string{"aaa", "aab", "aac", "aace"}
	// loop through each string and insert it into the tree
	for _, str := range s {
		t.Insert(str)
	}
	// print the root node of the tree with level 0
	Print(t.root, 0)
}

// Traverse traverses the tree by appending each node's value to the slice and printing it when reaching a leaf node
func Traverse(node *Node, s []rune) {
	// if node is nil, return
	if node == nil {
		return
	}
	// if node has a non-empty value, append it to the slice
	if node.value != 0 { // 0 is rune's zero value which means empty
		s = append(s, node.value)
	}
	// if node has no children, print the slice as a string
	if len(node.children) == 0 {
		fmt.Println(string(s))
		return
	}
	// loop through each child node and traverse it recursively with updated slice
	for _, child := range node.children {
		Traverse(child, s)
	}
}

func TestOut(t2 *testing.T) {
	// initialize an empty tree
	t := &Tree{}
	// insert some strings into the tree for testing
	t.Insert("aaa")
	t.Insert("aab")
	t.Insert("aac")
	t.Insert("aace")
	// traverse the tree with root node and empty slice
	Traverse(t.root, []rune{})
}
