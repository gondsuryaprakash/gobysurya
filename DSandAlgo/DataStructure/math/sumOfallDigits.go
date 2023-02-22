package math

func SumOfAllDigits(x int) int {

	var sum int

	for x > 0 {
		lastDigit := x % 10
		sum += lastDigit
		x = x / 10

	}
	return sum
}
