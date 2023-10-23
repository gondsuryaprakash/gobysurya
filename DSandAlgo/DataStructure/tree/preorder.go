package tree

import "fmt"

func PreOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Val)

	// DO recursive call for the left node
	PreOrder(node.Left)
	// DO recursive call for the right node
	PreOrder(node.Right)
}
