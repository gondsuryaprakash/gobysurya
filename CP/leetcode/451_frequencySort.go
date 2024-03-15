package leetcode

import "container/heap"

type characterFrequency struct {
	character string
	count     int
}

type frquencyCountHeap []characterFrequency

func (h frquencyCountHeap) Len() int {
	return len(h)
}
func (h frquencyCountHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h frquencyCountHeap) Less(i, j int) bool {
	return h[j].count < h[i].count
}

func (h *frquencyCountHeap) Push(x interface{}) {
	*h = append(*h, x.(characterFrequency))
}
func (h *frquencyCountHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func NewMaxCharacterHeap(arr []characterFrequency) *frquencyCountHeap {
	var minSoldersHeap = &frquencyCountHeap{}
	*minSoldersHeap = arr
	heap.Init(minSoldersHeap)
	return minSoldersHeap
}
func FrequencySort(s string) string {
	rc := make(map[string]int)
	for _, v := range s {
		vs := string(v)
		if _, ok := rc[vs]; !ok {
			rc[vs] = 1
		} else {
			rc[vs] += 1
		}
	}
	h := NewMaxCharacterHeap([]characterFrequency{})
	for k, v := range rc {
		heap.Push(h, characterFrequency{
			character: k,
			count:     v,
		})
	}

	result := ""
	for h.Len() > 0 {
		maxFreq := heap.Pop(h).(characterFrequency)
		k := maxFreq.count
		for k > 0 {
			result += maxFreq.character
			k--
		}
	}
	return result
}
