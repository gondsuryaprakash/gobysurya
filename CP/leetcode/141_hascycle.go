package leetcode

func HasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}

	return false
}
