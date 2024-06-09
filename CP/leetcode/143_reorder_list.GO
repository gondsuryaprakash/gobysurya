package leetcode

func ReorderList(head *ListNode) {

	arr := []int{}

	tmp := head

	for tmp.Next != nil {
		arr = append(arr, tmp.Val)
		tmp = tmp.Next
	}
	return nil
}
