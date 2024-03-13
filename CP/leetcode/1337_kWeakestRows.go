package leetcode

import "container/heap"

type Solders struct {
	Index int
	Count int
}

type soldersHeap []Solders

func (h soldersHeap) Len() int {
	return len(h)
}

func (h soldersHeap) Less(i, j int) bool {
	if h[i].Count == h[j].Count {
		return h[i].Index < h[j].Index
	}
	return h[i].Count < h[j].Count

}
func (h soldersHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *soldersHeap) Push(x interface{}) {
	*h = append(*h, x.(Solders))
}
func (h *soldersHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func NewMinSolders(arr []Solders) *soldersHeap {
	var minSoldersHeap = &soldersHeap{}
	*minSoldersHeap = arr
	heap.Init(minSoldersHeap)
	return minSoldersHeap
}

func KWeakestRows(mat [][]int, k int) []int {
	numberOfSolderes := []Solders{}
	for index, row := range mat {
		var solder Solders
		count := 0
		for _, v := range row {
			if v == 1 {
				count += 1
			}
		}
		solder.Count = count
		solder.Index = index
		numberOfSolderes = append(numberOfSolderes, solder)
	}

	minHeap := NewMinSolders(numberOfSolderes)

	result := []int{}
	for k > 0 {
		weakestRow := heap.Pop(minHeap)
		solder := weakestRow.(Solders)
		result = append(result, solder.Index)
		k--
	}
	return result
}
