package leetcode

import "github.com/spf13/cast"

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(elem string) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		lastIndex := len(*s) - 1
		lastEmen := (*s)[lastIndex]
		*s = (*s)[:lastIndex]
		return lastEmen
	}
}

func EvalRPN(tokens []string) int {

	var stack = &Stack{}

	for i := 0; i < len(tokens); i++ {
		// '+', '-', '*', and '/'.
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			stack.Push(tokens[i])
		} else {
			if tokens[i] == "+" {
				var sum int
				for len(*stack) > 0 {
					sum += cast.ToInt(stack.Pop())
				}
				stack.Push(string(sum))
			}
			if tokens[i] == "-" {
				var sum int
				for len(*stack) > 0 {
					sum -= cast.ToInt(stack.Pop())
				}
				stack.Push(string(sum))
			}
			if tokens[i] == "*" {
				var sum int
				for len(*stack) > 0 {
					sum *= cast.ToInt(stack.Pop())
				}
				stack.Push(string(sum))

			}
			if tokens[i] == "/" {
				var sum int
				for len(*stack) > 0 {
					sum /= cast.ToInt(stack.Pop())
				}
				stack.Push(string(sum))
			}
		}
	}

	return cast.ToInt((*stack)[0])

}
