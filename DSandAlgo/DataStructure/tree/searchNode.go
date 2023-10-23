package tree

// Geiven node search the node if exist yes else return false
func SearchNode(node *Node, key int) bool {

	if node == nil {
		return false
	}

	if node.Val == key {
		return true
	}
	if ok := SearchNode(node.Left, key); ok {
		return true
	}
	if ok := SearchNode(node.Right, key); ok {
		return true
	}
	return false
}
