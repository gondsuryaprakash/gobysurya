package leetcode

import (
	"container/heap"
	"sort"
)

type maxHeap []int

type MaxHeap interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]

}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func NewMaxHeap(arr []int) *maxHeap {
	var maxHeap = &maxHeap{}
	*maxHeap = arr
	heap.Init(maxHeap)
	return maxHeap
}

func MaxProduct(nums []int) int {

	maxHeap := NewMaxHeap(nums)
	a := heap.Pop(maxHeap)
	b := heap.Pop(maxHeap)
	return (a.(int) - 1) * (b.(int) - 1)
}

func DeleteGreatestValue(grid [][]int) int {

	rc := make(map[int]*maxHeap)
	for index, row := range grid {
		maxHeap := NewMaxHeap(row)
		rc[index] = maxHeap
	}

	result := 0
	for rc[0].Len() > 0 {

		max := 0
		for _, v := range rc {
			last := heap.Pop(v)
			if last.(int) > max {
				max = last.(int)
			}
		}
		result += max
	}

	return result
}
