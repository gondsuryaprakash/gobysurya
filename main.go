package main

import (
	"fmt"

	"gotutorial.com/CP/leetcode"
)

func main() {
	arr := []string{"2", "1", "+", "3", "*"}
	n := leetcode.EvalRPN(arr)
	fmt.Println(n)
}
