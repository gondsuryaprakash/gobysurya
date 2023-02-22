package math

func Palindrom(x int) bool {

	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	reveres := ReverseNumber(x)

	return reveres == x
}
