package concurrency

import "fmt"

func Common(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d run : ", i)
	}
}

func PrintTo15() {
	Common(15)
}
func PrintTo10() {
	Common(10)
}
