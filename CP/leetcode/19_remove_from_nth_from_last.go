package leetcode

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	start := head
	fast := start
	slow := start

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return start
}
