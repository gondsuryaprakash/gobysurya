package math

func ReverseNumber(x int) int {
	var reverse int
	for x > 0 {

		lastDigit := x % 10
		reverse = reverse*10 + lastDigit
		x = x / 10

	}
	return reverse
}
