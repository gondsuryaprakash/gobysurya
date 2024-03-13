package leetcode

import (
	"container/heap"
	"math"
)

type maxHeap1 []int

func (h maxHeap1) Len() int {
	return len(h)
}
func (h maxHeap1) Less(i, j int) bool {
	return h[i] > h[j]

}
func (h maxHeap1) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap1) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func NewMaxHeap1(arr []int) *maxHeap1 {
	var maxHeap = &maxHeap1{}
	*maxHeap = arr
	heap.Init(maxHeap)
	return maxHeap
}

func pickGifts(gifts []int, k int) int64 {
	h := NewMaxHeap1(gifts)
	for k > 0 {
		gift := heap.Pop(h).(int)
		floatMathSqrt := math.Floor(math.Sqrt(float64(gift)))
		heap.Push(h, int(floatMathSqrt))
		k--
	}
	var sum int64
	for h.Len() > 0 {
		sum += int64(heap.Pop(h).(int))
	}
	return sum
}
