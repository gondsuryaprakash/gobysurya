package main

func SwapNumber(a, b int) (int, int) {
	a = a - b
	b = a + b
	a = b - a
	return a, b
}
