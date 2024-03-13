package leetcode

import (
	"container/heap"
	"strconv"
)

type Player struct {
	Score int
	Index int
}

type RankList []Player

func (h RankList) Len() int {
	return len(h)
}
func (h RankList) Less(i, j int) bool {
	return h[i].Score > h[j].Score

}
func (h RankList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *RankList) Push(x interface{}) {
	*h = append(*h, x.(Player))
}
func (h *RankList) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func NewPlayerHeap(arr []Player) *RankList {
	var maxHeap = &RankList{}
	*maxHeap = arr
	heap.Init(maxHeap)
	return maxHeap
}

func FindRelativeRanks(score []int) []string {

	rankArr := []Player{}
	for index, v := range score {
		rankArr = append(rankArr, Player{
			Score: v,
			Index: index,
		})
	}

	result := []string{}

	for i := 0; i < len(score); i++ {
		result = append(result, "")
	}
	rc := map[int]string{
		0: "Gold Medal",
		1: "Silver Medal",
		2: "Bronze Medal",
	}

	h := NewPlayerHeap(rankArr)
	for i := 0; i < len(score); i++ {
		// if h.Len() > 1 {
		topPlayer := heap.Pop(h).(Player)
		v, ok := rc[i]
		if ok {
			result[topPlayer.Index] = v
		} else {
			result[topPlayer.Index] = strconv.Itoa(i + 1)
		}

		// }
	}
	return result
}
