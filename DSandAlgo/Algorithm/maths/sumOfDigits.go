package maths

func SumOfNumber(n int) int {

	if n < 10 {
		return n
	} else {
		return sum(n)
	}
}

func sum(n int) int {
	s := 0

	for n > 0 {
		s = s + n%10
		n = n / 10
	}
	return s
}
