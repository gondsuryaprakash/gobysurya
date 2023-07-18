package leetcode

import "math"

type StackStock struct {
	arr []int
}

func (s *StackStock) Push(ele int) {
	s.arr = append(s.arr, ele)
}



func (s *StackStock) Top() int {
	if len(s.arr) > 0 {
		return s.arr[len(s.arr)-1]
	}
	return math.MaxInt32
}

func MaxProfit(prices []int) int {
	stack := StackStock{
		arr:    []int{},
		
	}
	var maxProfit int
	stack.Push(prices[0])
	for i := 1; i < len(prices); i++ {
		if stack.Top() > prices[i] {
			stack.Push(prices[i])
		} else {
			newProfit := prices[i] - stack.Top()
			if maxProfit < newProfit {
				maxProfit = newProfit
			}
		}
	}
	return maxProfit
}
