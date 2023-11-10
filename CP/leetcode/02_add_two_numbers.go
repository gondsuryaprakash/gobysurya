package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dump := &ListNode{}
	tmp := dump

	var carry int
	for l1 != nil || l2 != nil || carry > 0 {
		s := 0
		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}
		s += carry
		carry = s / 10
		tmp.Next = &ListNode{
			Val: s,
		}
		tmp = tmp.Next
	}
	return dump.Next

}
