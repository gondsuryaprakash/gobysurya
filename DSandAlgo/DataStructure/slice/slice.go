package slice

import "fmt"

func Change1(abc []int) {

	for i := range abc {
		abc[i] = 5
	}
	fmt.Println(abc)
}

func MainFunctionSlice1() {
	abc := []int{1, 2, 3}

	Change1(abc)

	fmt.Println(abc)
}

func Change2(abc []int) {
	abc = append(abc, 5)
	for i := range abc {
		abc[i] = 5
	}

	fmt.Println(abc)
}

func MainFunctionSlice2() {
	abc := []int{1, 2, 3}
	Change2(abc)
	fmt.Println(abc)
}

func Change3(abc []int) []int {
	abc = append(abc, 5)
	for i := range abc {
		abc[i] = 5
	}

	fmt.Println(abc)
	return abc
}

func MainFunctionSlice3() {
	abc := []int{1, 2, 3}
	abc = Change3(abc)
	fmt.Println(abc)
}


