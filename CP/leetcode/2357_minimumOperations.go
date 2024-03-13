package leetcode

import (
	"container/heap"
	"sort"
)

type minHeap2357 []int

type MinHeap2357 interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func (h *minHeap2357) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap2357) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h minHeap2357) Len() int {
	return len(h)
}

func (h minHeap2357) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap2357) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func NewMinHeap(arr []int) *maxHeap {
	var maxHeap = &maxHeap{}
	*maxHeap = arr
	heap.Init(maxHeap)
	return maxHeap
}

func minimumOperations(nums []int) int {
	// remove all the zero
	newArray := []int{}
	for _, v := range nums {
		if v == 0 {
			newArray = append(newArray, v)
		}
	}
	// heap := NewMinHeap(newArray)

	return 0
}
