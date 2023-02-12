package main

import (
	"fmt"

	search "gotutorial.com/DSandAlgo/Algorithm/Search"
)

func main() {
	arr := []int{19, 10, 5, 6}
	i, err := search.LinearSearch(5, arr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Element found at index : ", i)
	}

}
