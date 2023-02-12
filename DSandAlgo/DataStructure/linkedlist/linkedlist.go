package linkedlist

import "fmt"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	val  int
	next *Node
}

// First return head of the linkedList
func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

// Add node in LinkedList
func (l *List) Push(val int) {
	node := &Node{val: val}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *List) AddAtFirst(val int) {

	n := &Node{val: val}
	n.next = l.head
	l.head = n

}

func (l *List) AddInAfterNode(val int, prev *Node) {

	if prev == nil {
		return
	}
	n := &Node{val: val}
	n.next = prev.next
	prev.next = n

}
func (l *List) PrintLinkedList() {
	n := l.First()
	for {
		fmt.Println(n.val)

		n = n.Next()
		if n == nil {
			break
		}
	}
}

func (l *List) SearchNode(key int) bool {

	first := l.First()

	for {

		if first.val == key {
			return true
		}
		first = first.next

		if first.next == nil {
			break
		}
	}

	return false

}

func (l *List) InsertAtLast(val int) {
	n := &Node{val: val}

	if l.head == nil {
		l.head = n
	}

	h := l.First()
	for {

		if h.next == nil {
			break
		}

		h = h.next
	}

	h.next = n

}

// This function reverse the linked list
func (l *List) ReverseLinkedList() {

	head := l.First()
	prev := &Node{}
	next := &Node{}
	for head != nil {
		next = head.next
		head.next = prev
		prev = head
		head = next

	}
	l.head = prev

}

func (l *List) FindMiddle() {
	head := l.First()

	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

	}

	fmt.Println("Middle value of the LinkedList : ", slow.val)

}

func LinkedList() {

	l := &List{}

	l.Push(1)
	l.Push(2)
	l.Push(3)

	// Before the Add in front
	l.PrintLinkedList()

	last := l.Last()
	fmt.Println("Last.val ", last.val)

	// After add in front
	l.AddAtFirst(5)
	l.PrintLinkedList()
	l.AddInAfterNode(28, l.head.next)

	fmt.Println("Afer Operation")
	l.PrintLinkedList()

	fmt.Println("Insert at last : ")
	l.InsertAtLast(2229)
	l.PrintLinkedList()

	// l.ReverseLinkedList()
	fmt.Println("Reverse LinkedList : ")
	l.PrintLinkedList()
	l.FindMiddle()

	// isNodePresent := l.SearchNode(281)

	// fmt.Println("isNodePresent ", isNodePresent)

}
