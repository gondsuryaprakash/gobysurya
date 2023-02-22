package stack

import "fmt"

type stack []int

// Check for Empty
func (s stack) Empty() bool {
	return len(s) == 0
}

// Push the value in stack
func (s *stack) Push(v int) {
	*s = append(*s, v)
}

// Pop the value in stack and return top of the element
func (s *stack) Pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func CallStack() {
	s := stack{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
}
