package leetcode

import "container/heap"

type kthSmallestHeap []int

func (h kthSmallestHeap) Len() int {
	return len(h)
}
func (h kthSmallestHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h kthSmallestHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *kthSmallestHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *kthSmallestHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func NewkthSmallestHeap(arr []int) *kthSmallestHeap {
	newArray := &kthSmallestHeap{}
	*newArray = arr
	heap.Init(newArray)
	return newArray
}

func kthSmallest(matrix [][]int, k int) int {

	h := NewkthSmallestHeap([]int{})

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			heap.Push(h, matrix[i][j])
		}
	}

	var p int
	for k > 0 {
		p = heap.Pop(h).(int)
		k--
	}

	return p
}
