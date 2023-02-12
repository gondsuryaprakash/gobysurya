package bst

import "fmt"

type Tree struct {
	node *Node
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (bst *Tree) Insert(val int) *Tree {
	if bst.node == nil {
		bst.node = &Node{Val: val}
	} else {
		bst.node.Insert(val)
	}
	return bst
}

func (node *Node) Insert(val int) {
	if val <= node.Val {
		if node.Left == nil {
			node.Left = &Node{Val: val}
		} else {
			node.Left.Insert(val)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Val: val}
		} else {
			node.Right.Insert(val)
		}
	}

}
func (n *Node) IsExist(val int) bool {
	if n == nil {
		return false
	}

	if n.Val == val {
		return true
	}
	if val <= n.Val {
		return n.Left.IsExist(val)
	} else {
		return n.Right.IsExist(val)
	}

}

func PrintNode(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.Val)
	PrintNode(n.Left)
	PrintNode(n.Right)

}

func BST() {
	t := &Tree{}

	t.Insert(10).
		Insert(8).
		Insert(20).
		Insert(9).
		Insert(0).
		Insert(15).
		Insert(25)
	PrintNode(t.node)

	isPresent := t.node.IsExist(110)
	fmt.Println("Node is present ", isPresent)

}
