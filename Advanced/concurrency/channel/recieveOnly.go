package channel

import (
	"fmt"
	"math/rand"
	"time"
)

func LongTimeRequest() <-chan int32 {

	ch := make(chan int32)

	go func() {
		time.Sleep(time.Second * 3)
		ch <- rand.Int31()
	}()
	return ch

}

func Squares(a, b int32) int32 {
	return a*a + b*b
}

func CallLongTimeRequest() {
	a := LongTimeRequest()
	b := LongTimeRequest()

	fmt.Println(Squares(<-a, <-b))
}
