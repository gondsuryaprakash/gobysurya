package tree

import "fmt"

func PostOrder(node *Node) {
	if node == nil {
		return
	}
	// DO recursive call for the left node
	PostOrder(node.Left)
	// DO recursive call for the right node
	PostOrder(node.Right)

	fmt.Println(node.Val)
}
