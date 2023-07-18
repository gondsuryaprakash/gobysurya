package promises

import (
	"fmt"
	"math/rand"
	"time"
)

func SumOfSquareSquare(a, b int) int {
	return a*a + b*b
}

// This function will take time upto 3s
func longTimeRequestRequest() <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		ch <- rand.Intn(100)
	}()
	return ch
}

func CallPromiseFunction() {
	a, b := longTimeRequestRequest(), longTimeRequestRequest()

	fmt.Println(SumOfSquareSquare(<-a, <-b))

}
