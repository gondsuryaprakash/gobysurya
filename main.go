package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"gotutorial.com/CP/leetcode"
)

func main() {
	// rect := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	area := leetcode.DetectCapitalUse("aLL")

	fmt.Println(area)
}
