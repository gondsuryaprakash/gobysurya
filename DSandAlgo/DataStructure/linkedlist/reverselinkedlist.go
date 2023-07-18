package linkedlist

import "fmt"

type ListNode struct {
	val  int
	Next *ListNode
}

type SilgleList struct {
	Head *ListNode
	Len  int
}

func Init() *ListNode {
	return &ListNode{}
}

func (l *ListNode) Traverse() {
	for l != nil {
		fmt.Printf("Val of the list : %d", l.val)
		l = l.Next
	}
}

func (s *SilgleList) Reverse() {
	curr := s.Head
	var prev *ListNode
	var next *ListNode

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	s.Head = prev

}

func (s *SilgleList) AddFront(val int) {
	ele := &ListNode{
		val: val,
	}
	if s.Head == nil {
		s.Head = ele
	} else {
		ele.Next = s.Head
		s.Head = ele
	}

}
