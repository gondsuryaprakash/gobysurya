package leetcode

func RotateRight(head *ListNode, k int) *ListNode {

	if head == nil || k == 0 {
		return nil
	}
	length := 0

	curr := head
	dummy := &ListNode{Next: head}

	for curr != nil {
		length++
		curr = curr.Next
	}
	toRotate := length - (k % length)

	tmp := head
	for toRotate > 1 {
		toRotate--
		tmp = tmp.Next
	}

	curr.Next = head
	dummy.Next = tmp.Next
	tmp.Next = nil

	return dummy.Next

}
