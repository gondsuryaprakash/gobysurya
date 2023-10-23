package tree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

func CallNode() {

	newNode := NewNode(1)
	newNode.Left = NewNode(2)
	newNode.Left.Left = NewNode(3)
	newNode.Left.Right = NewNode(4)
	newNode.Right = NewNode(10)
	newNode.Right.Left = NewNode(6)
	newNode.Right.Right = NewNode(7)

	InOrder(newNode)
}

func InOrder(node *Node) {

	if node == nil {
		return
	}

	// Do recursive call for the left node
	InOrder(node.Left)

	fmt.Println(node.Val)
	// Do recursive call for the right node
	InOrder(node.Right)

}




