package channel

import (
	"fmt"
	"time"
)

func BlockChannel() {

	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	val := <-ch

	// I want to do another task
	fmt.Println("another task I have done !!")

	fmt.Println("Value Recived :", val)
}

// to solve above bloacking issue what can we do we use select statment

func NonBlockingChannel() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()

	for {
		select {
		case val := <-ch:
			fmt.Println("Value recieved : ", val)
			return
		default:
			fmt.Println("no value recieved !")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
