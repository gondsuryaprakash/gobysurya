package contest

import "strconv"

func sumOfEncryptedInt(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += greatestNumberInInteger(nums[i])
	}

	return sum
}

func greatestNumberInInteger(x int) int {
	stringX := strconv.Itoa(x)
	var maxNumber int
	for x > 0 {
		reminder := x % 10
		maxNumber = isMax(maxNumber, reminder)
		x = x / 10
	}
	start := 0
	s := ""
	for start < len(stringX) {
		s += strconv.Itoa(maxNumber)
		start++
	}

	num, _ := strconv.Atoi(s)
	return num
}

func isMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
