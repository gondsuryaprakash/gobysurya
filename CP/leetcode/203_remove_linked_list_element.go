package leetcode

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func RemoveElements(head *ListNode2, val int) *ListNode2 {

	if head == nil {
		return nil
	}

	dummy := &ListNode2{Next: head}
	prev, curr := dummy, head

	for curr != nil {
		next := curr.Next
		if curr.Val == val {
			prev.Next = next
		} else {
			prev = curr
		}

		curr = next
	}

	return dummy.Next
}
