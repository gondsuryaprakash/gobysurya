package promises

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest(ch chan<- int) {
	time.Sleep(time.Second * 3)
	a := rand.Intn(100)
	ch <- a
}

func CallPromiseWithArgs() {
	args := make(chan int, 2)
	go longTimeRequest(args)
	go longTimeRequest(args)
	fmt.Println(SumOfSquareSquare(<-args, <-args))
}
