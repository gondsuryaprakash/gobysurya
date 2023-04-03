package channel

import (
	"fmt"
	"time"
)

func SumWithGoRoutine() {
	res := make(chan int)
	a := 5
	b := 10
	go Sum(a, b, res)
	sum := <-res
	fmt.Println(sum)
	time.Sleep(2 * time.Second)
}



func Sum(a, b int, ch chan<- int) {
	sum := a + b
	ch <- sum
}
