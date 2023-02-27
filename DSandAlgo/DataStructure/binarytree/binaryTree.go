package binarytree

import "fmt"

// node struct
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Binary Tree struct
type BinaryTree struct {
	Root *Node
}

// Insert in Binary Tree
func (b *BinaryTree) AddNode(val int) *BinaryTree {

	// Creating a new node
	node := &Node{
		Val: val,
	}

	// Check b.Root is nil
	if b.Root == nil {
		b.Root = node
		return b
	} else {
		focusNode := b.Root
		var parentNode *Node
		for {
			parentNode = focusNode

			if val < focusNode.Val {
				focusNode = focusNode.Left
				if focusNode == nil {
					parentNode.Left = node
					return b
				}
			} else {
				focusNode = focusNode.Right
				if focusNode == nil {
					parentNode.Right = node
					return b
				}
			}
		}

	}

}

// PreOrder Traversal of Binary Tree
func (b *BinaryTree) PreOrderTraversal(root *Node) {
	if root != nil {
		fmt.Println(root.Val, " ", "--->")
		b.PreOrderTraversal(root.Left)
		b.PreOrderTraversal(root.Right)
	}
}

func CallBinaryTree() {

	bt := &BinaryTree{}

	bt.AddNode(51)

	bt.AddNode(26)

	bt.AddNode(76)

	bt.AddNode(13)

	bt.AddNode(38)

	bt.AddNode(44)

	bt.AddNode(31)
	bt.PreOrderTraversal(bt.Root)
}
