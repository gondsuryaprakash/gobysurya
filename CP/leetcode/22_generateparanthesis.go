package leetcode

import "fmt"

var record = map[string]string{}

func GenerateParanthesis(n int) []string {
	parenthesis := ""
	for i := 0; i < n; i++ {
		parenthesis += "()"
	}
	arrParanthesis := []string{}
	Generator(parenthesis, "", &arrParanthesis)
	return arrParanthesis
}

func Generator(str string, answer string, result *[]string) {
	if len(str) == 0 && isValidParenthesis(answer) {
		*result = append(*result, answer)
	}
	for i := 0; i < len(str); i++ {
		ch := string(str[i])
		left := str[:i]
		right := str[i+1:]
		rest := left + right
		if !IsDuplicate(answer + ch) {
			Generator(rest, answer+ch, result)
		}
	}
}

func isValidParenthesis(str string) bool {
	s := stack{}
	for i := 0; i < len(str); i++ {

		if string(str[i]) == "(" {
			s.Push("(")
		} else {
			if len(s) > 0 {
				if s.Head() == "(" {
					s.Pop()
				} else {
					return false
				}
			}
		}
	}
	return s.IsEmpty()
}

type stack []string

func (s *stack) Head() string {
	return (*s)[len(*s)-1]
}
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *stack) Push(str string) {
	*s = append(*s, str)
}
func (s *stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	lastElement := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return lastElement
}
func IsDuplicate(str string) bool {
	if record == nil {
		fmt.Println("Nil")
	}
	if _, ok := record[str]; ok {
		return true
	}
	record[str] = ""
	return false
}
