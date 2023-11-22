package leetcode

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	dummyA := headA
	dummyB := headB
	for dummyA != dummyB {
		if dummyA == nil {
			dummyA = headB
		} else {
			dummyA = dummyA.Next
		}
		if dummyB == nil {
			dummyB = headA
		} else {
			dummyB = dummyB.Next
		}
	}
	return dummyA
}
