package leetcode

import (
	"container/heap"
)

type Pairs struct {
	sum int
	i   int
	j   int
}

type minPairsHeap []Pairs

func (h minPairsHeap) Len() int {
	return len(h)
}
func (h minPairsHeap) Less(i, j int) bool {
	return h[i].sum < h[j].sum

}
func (h minPairsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minPairsHeap) Push(x interface{}) {
	*h = append(*h, x.(Pairs))
}
func (h *minPairsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func NewMinPairsHeap(arr []Pairs) *minPairsHeap {
	var minHeap = &minPairsHeap{}
	*minHeap = arr
	heap.Init(minHeap)
	return minHeap
}

func KSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	p := []Pairs{}
	for i := 0; i < k; i++ {
		p = append(p, Pairs{
			sum: nums1[i] + nums2[0],
			i:   i,
			j:   0,
		})
	}
	result := [][]int{}
	h := NewMinPairsHeap(p)

	for h.Len() > 0 && k > 0 {
		currentPair := heap.Pop(h).(Pairs)
		result = append(result, []int{nums1[currentPair.i], nums2[currentPair.j]})
		nextElement := currentPair.j + 1
		if nextElement < len(nums2) {
			heap.Push(h, Pairs{sum: nums1[currentPair.i] + nums2[nextElement],
				i: currentPair.i,
				j: nextElement})
		}
		k--
	}
	return result
}
