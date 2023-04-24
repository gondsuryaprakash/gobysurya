package closure

func Adder() func(int) int {

	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}
