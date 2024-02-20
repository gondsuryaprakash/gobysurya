package recursion

func Fibnacci(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return Fibnacci(n-1) + Fibnacci(n+2)
}
