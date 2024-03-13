package leetcode

import (
	"fmt"
)

func RemoveZeroSumSublists(head *ListNode) *ListNode {

	array := []int{}
	curr := head

	mp := make(map[int]ListNode)

	prefix := 0
	mp[prefix] = ListNode{
		Val:  0,
		Next: head,
	}

	for curr.Next != nil {
		prefix += curr.Val
		mp[prefix] = *curr
		curr = curr.Next
	}

	for head != nil {

	}

	fmt.Println(array)
	return nil
}
