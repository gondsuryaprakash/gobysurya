package channel

import (
	"fmt"
	"time"
)

func fast(num int, out chan<- int) {

	product := num * 2

	out <- product
}

func slow(num int, out chan<- int) {

	product := num * 2

	time.Sleep(15 * time.Microsecond)

	// Sending data
	out <- product
}

func CallSlowAndFastChannel() {

	out1 := make(chan int)
	out2 := make(chan int)

	go fast(5, out1)
	go slow(6, out2)

	select {
	case res:= <-out1:
		fmt.Println("Fast Finished First", res)
	case res1:= <-out2:
		fmt.Println("Fast Finished First", res1)
	}
}

