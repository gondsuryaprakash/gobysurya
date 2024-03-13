package leetcode

import "container/heap"

func LastStoneWeight(stones []int) int {

	h := NewMaxHeap(stones)
	for h.Len() >= 2 {
		maxStone1 := heap.Pop(h).(int)
		maxStone2 := heap.Pop(h).(int)
		if maxStone1 != maxStone2 {
			heap.Push(h, maxStone1-maxStone2)
		}
	}
	if h.Len() == 1 {
		return heap.Pop(h).(int)

	}
	return 0
}
