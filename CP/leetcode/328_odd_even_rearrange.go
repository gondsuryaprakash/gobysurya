package leetcode

func OddEvenList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	if head.Next == nil || head.Next.Next == nil {
		return head
	}
	odd_head := head
	even_head := head.Next
	odd := odd_head
	even := even_head

	for even != nil && even.Next != nil {

		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}

	odd.Next = even_head

	return head

}
