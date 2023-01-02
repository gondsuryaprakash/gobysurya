package linkedlist

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Append(node *Node) {
	second := l.head
	l.head = node
	l.head.Next = second
	l.length++

}
