package main

import (
	"fmt"

	"gotutorial.com/DSandAlgo/Algorithm/sort"
)

func main() {
	arr := []int{64, 25, 12, 22, 11}
	arr1 := sort.BubbleSort(arr)
	fmt.Println(arr1)
}
