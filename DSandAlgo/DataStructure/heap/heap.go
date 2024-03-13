package heap

import (
	"container/heap"
	"fmt"
	"sort"
)

// type MaxHeap struct {
// 	array []int
// }

// func (h *MaxHeap) Insert(key int) {
// 	h.array = append(h.array, key)
// 	h.MaxHeapifyUp(len(h.array) - 1)
// }

// func (h *MaxHeap) MaxHeapifyUp(index int) {

// }

// func Parent(i int) int {
// 	return (i - 1) / 2
// }

// func Left(i int) int {
// 	return 2*i + 1
// }

// func Right(i int) int {
// 	return 2*i + 2
// }

// func (h *MaxHeap) Swap(i1, i2 int) {
// 	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
// }

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
	fmt.Println("New MinHeap Created")
	return minHeap
}

func NumberGame(nums []int) []int {
	result := []int{}
	minHeap := NewHeap(nums)
	for minHeap.Len() > 0 {
		alice := heap.Pop(minHeap)
		bob := heap.Pop(minHeap)
		result = append(result, bob.(int))
		result = append(result, alice.(int))
	}

	return result
}
