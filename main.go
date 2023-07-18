package main

import (
	_ "net/http/pprof"

	"gotutorial.com/Advanced/regularexpression"
)

func main() {

	// arr := []int{2, 2, 5, 10, 1, 2, 10, 5, 10, 10, 2}
	// arr := []int{2, 3, 1, 1, 4}

	regularexpression.IsMobile("ios")

}
