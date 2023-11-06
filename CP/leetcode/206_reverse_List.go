package leetcode

type ListNode3 struct {
	Val  int
	Next *ListNode3
}

func ReverseList(head *ListNode3) *ListNode3 {
	if head == nil {
		return nil
	}

	curr := head
	var prev *ListNode3
	var next *ListNode3

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev

}
