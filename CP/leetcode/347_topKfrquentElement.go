package leetcode

import "container/heap"

type frquency struct {
	count   int
	element int
}

type topKFrequencyHeap []frquency

func (h topKFrequencyHeap) Len() int {
	return len(h)

}
func (h topKFrequencyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h topKFrequencyHeap) Less(i, j int) bool {
	return h[j].count < h[i].count
}
func (h *topKFrequencyHeap) Push(x interface{}) {
	*h = append(*h, x.(frquency))
}
func (h *topKFrequencyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func NewTopKFrequencyHeap(arr []frquency) *topKFrequencyHeap {
	newArray := &topKFrequencyHeap{}
	*newArray = arr

	heap.Init(newArray)
	return newArray
}
func topKFrequent(nums []int, k int) []int {

	h := NewTopKFrequencyHeap([]frquency{})
	rc := make(map[int]int)

	for _, v := range nums {
		rc[v] += 1
	}

	for k, v := range rc {
		heap.Push(h, frquency{
			count:   v,
			element: k,
		})
	}

	result := []int{}

	for k > 0 {
		top := heap.Pop(h).(frquency)
		result = append(result, top.element)
		k--
	}
	return result
}
