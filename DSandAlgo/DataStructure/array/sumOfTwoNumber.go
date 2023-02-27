package array

func SumOfTwo() []int {
	arr := []int{3, 2, 4}
	target := 6
	maps := make(map[int]int, len(arr))

	for i, num2 := range arr {

		num1 := target - num2
		if j, ok := maps[num1]; ok {
			return []int{i, j}

		} else {
			maps[num2] = j
		}
	}
	return nil
}
