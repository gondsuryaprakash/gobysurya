package array

import "fmt"

func LeftRotate(arr []int, d int) {

	p := 1

	for p <= d {

		last := arr[0]

		for i := 0; i < len(arr)-1; i++ {
			arr[i] = arr[i+1]
		}

		arr[len(arr)-1] = last
		p += 1
	}

	fmt.Println(arr)
}
