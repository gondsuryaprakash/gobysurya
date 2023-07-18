package leetcode

import "math"

type StackStock1 struct {
	arr []int
}

func (s *StackStock1) Push(ele int) {
	s.arr = append(s.arr, ele)
}

func (s *StackStock1) Top() int {
	if len(s.arr) > 0 {
		return s.arr[len(s.arr)-1]
	}
	return math.MaxInt32
}

func (s *StackStock1) Pop() int {
	var lastElement int = math.MaxInt32
	if len(s.arr) > 0 {
		lastElement = s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
	}
	return lastElement
}

func (s *StackStock1) IsEmpty() bool {
	return len(s.arr) == 0
}

func MaxProfit1(prices []int) int {
	stack := StackStock1{
		arr: []int{},
	}
	var maxProfit int
	var totalProfit int

	for i := 0; i < len(prices); i++ {
		if stack.IsEmpty() {
			stack.Push(prices[i])
		}
		if stack.Top() > prices[i] {
			if stack.IsEmpty() {
				stack.Pop()
			}
			stack.Push(prices[i])
		} else {
			newProfit := prices[i] - stack.Top()
			if maxProfit < newProfit {
				maxProfit = 0
				totalProfit += newProfit
				
			}
		}
	}
	return totalProfit
}
