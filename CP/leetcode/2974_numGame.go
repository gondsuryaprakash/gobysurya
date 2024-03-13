package leetcode

import (
	"container/heap"
	"sort"
)

type minHeap []int

type Heap interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]

}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func NewHeap(arr []int) *minHeap {
	var minHeap = &minHeap{}
	*minHeap = arr
	heap.Init(minHeap)
	return minHeap
}

// func NumberGame(nums []int) []int {
// 	result := []int{}
// 	minHeap := NewHeap(nums)
// 	for minHeap.Len() > 0 {
// 		alice := heap.Pop(minHeap)
// 		bob := heap.Pop(minHeap)
// 		result = append(result, bob.(int))
// 		result = append(result, alice.(int))
// 	}

// 	return result
// }
