package channel

import (
	"fmt"
	"math/rand"
	"time"
)

func LongRequestSendOnly(ch chan<- int32) {
	time.Sleep(time.Second)
	ch <- rand.Int31n(100)
}

func CallLongTimeRequestSendOnly() {

	rand.Seed(time.Now().UnixNano())
	a, b := make(chan int32), make(chan int32)

	go LongRequestSendOnly(a)
	go LongRequestSendOnly(b)

	sum := Squares(<-a, <-b)

	fmt.Println(sum)

}
