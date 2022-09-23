package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Node :represents the components of a binary search tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BST : represents the binary search tree structure
type BST struct {
	Root *Node
	Len  int
}

func (n Node) String() string {
	return strconv.Itoa(n.Value)
}

func (b BST) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b BST) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.Root)
}

func (b BST) inOrderTraversalByNode(sb *strings.Builder, root *Node) {
	if root == nil {
		return
	}

	b.inOrderTraversalByNode(sb, root.Left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	b.inOrderTraversalByNode(sb, root.Right)
}

// Insert : will add anode to the tree
// the key to add should not be already in the tree
func (b *BST) Insert(value int) {
	b.Root = b.insertByNode(b.Root, value)
	b.Len++
}

func (b *BST) insertByNode(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}

	if value < root.Value {
		root.Left = b.insertByNode(root.Left, value)
	} else {
		root.Right = b.insertByNode(root.Right, value)
	}

	return root
}

// Search : will take in a key value
// and return true if there is a node with that value
func (b *BST) Search(value int) (*Node, bool) {
	return b.searchByNode(b.Root, value)
}

func (b *BST) searchByNode(root *Node, value int) (*Node, bool) {
	if root == nil {
		return nil, false
	}

	if value == root.Value {
		return root, true
	} else if value < root.Value {
		return b.searchByNode(root.Left, value)
	} else {
		return b.searchByNode(root.Right, value)
	}
}

// Remove : will take a key value
// and delete it from the BST
func (b *BST) Remove(value int) {
	b.removeByNode(b.Root, value)
}

func (b *BST) removeByNode(root *Node, value int) *Node {
	if root == nil {
		return root
	}

	if value > root.Value {
		root.Right = b.removeByNode(root.Right, value)
	} else if value < root.Value {
		root.Left = b.removeByNode(root.Left, value)
	} else {
		if root.Left == nil {
			return root.Right
		} else {
			tmp := root.Left
			for tmp.Right != nil {
				tmp = tmp.Right
			}

			root.Value = tmp.Value
			root.Left = b.removeByNode(root.Left, value)
		}
	}

	return root
}

func main() {
	n := &Node{1, nil, nil}
	n.Left = &Node{0, nil, nil}
	n.Right = &Node{2, nil, nil}
	b := BST{n, 3}
	fmt.Println(b)

	b.Insert(4)
	b.Insert(5)
	b.Insert(6)
	fmt.Println(b)

	fmt.Println(b.Search(6))
	fmt.Println(b.Search(3))

	fmt.Printf("\n%v\n", b)
	b.Remove(6)
	fmt.Println(b)
	fmt.Println(b.Search(6))
}
