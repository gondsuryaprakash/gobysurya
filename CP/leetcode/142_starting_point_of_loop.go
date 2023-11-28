package leetcode

func DetectCycle(head *ListNode) *ListNode {

	if head == nil && head.Next == nil {
		return nil
	}

	slow := head
	fast := head
	entryPoint := head

	for fast != nil && fast.Next != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {

			for entryPoint != fast {
				entryPoint = entryPoint.Next
				fast = fast.Next

			}
			return entryPoint
		}
	}

	return nil
}
