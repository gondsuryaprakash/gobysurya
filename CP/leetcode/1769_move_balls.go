package leetcode

func MinOperations(boxes string) []int {

	n := len(boxes)

	ans := make([]int, n)

	op, ball := 0, 0
	for i := 0; i < len(boxes); i++ {
		ans[i] += op
		if boxes[i] == '1' {
			ball++
		}
		op += ball
	}

	op, ball = 0, 0
	for i := n - 1; i >= 0; i-- {
		ans[i] += op
		if boxes[i] == '1' {
			ball++
		}
		op += ball
	}

	return ans
}
