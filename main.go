package main

import (
	"fmt"

	"gotutorial.com/DSandAlgo/DataStructure/array"
)

func main() {
	arr := []int{12, 13, 1, 10, 34, 1}
	third := array.FindThirdHighestNumber(arr)
	fmt.Println(third)

}
