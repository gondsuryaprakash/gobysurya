package main

import (
	"fmt"
	"strings"
)

func main() {

	name1 := "Surya"
	name2 := "Surya Prakash Gond"

	fmt.Println(strings.Compare(name1, "Surya Prakash"))
	fmt.Println(strings.Compare(name2, "Surya Prakash"))

}
