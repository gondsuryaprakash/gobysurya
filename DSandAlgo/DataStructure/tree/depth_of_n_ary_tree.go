package tree

import "math"

/*Given an N-Ary tree, find depth of the tree. An N-Ary tree is a tree in which nodes can have at most N children. - GeeksForGeeks*/

type NAryNode struct {
	Key   int
	Child []NAryNode
}

func newNAryNode(key int) *NAryNode {
	return &NAryNode{
		Key: key,
	}
}

func DepthOfNAryTree(node *NAryNode) int {

	if node == nil {
		return 0
	}

	maxDepth := 0

	for i, val := range node.Child {

		maxDepth = int(math.Max(float64(maxDepth), float64(DepthOfNAryTree(&val.Child[i]))))
	}

	return maxDepth + 1
}
