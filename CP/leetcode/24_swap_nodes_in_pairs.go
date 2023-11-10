package leetcode

func SwapPairs(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	curr := head
	dummy := &ListNode{Val: -1}
	prev := dummy
	dummy.Next = head
	for curr != nil && curr.Next != nil {
		prev.Next = curr.Next
		curr.Next = curr.Next.Next
		prev.Next.Next = curr
		curr = curr.Next
		prev = prev.Next.Next
	}
	return dummy.Next

}
