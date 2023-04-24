package channel

import (
	"fmt"
	"math/rand"
	"time"
)

func Source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1

	//Sleep for 1s/2s/3s
	time.Sleep(time.Duration(rb) * time.Second)

	c <- ra
}

func CallSource() {
	rand.Seed(time.Now().UnixNano())

	stratTime := time.Now()
	ch := make(chan int32, 5)

	for i := 0; i < cap(ch); i++ {
		go Source(ch)
	}

	rand := <-ch

	fmt.Println(time.Since(stratTime))
	fmt.Println(rand)
}
