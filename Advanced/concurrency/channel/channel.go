package channel

import (
	"fmt"
	"time"
)

func CreateChannel() {

	message := make(chan string)

	go func() {
		message <- "From channel"
	}()
	fmt.Println("Waiting for Goroutine")
	fmt.Print(<-message)
}

func CreateAnotherChannel(msg chan<- int) {

	if true {
		msg <- 12
	}

}

func Channel1() {
	fmt.Println("Channel1")
	ch := make(chan int)
	go send(ch)
	go recive(ch)
	time.Sleep(time.Second * 1)
}

func send(ch chan int) {
	fmt.Println("In Send")
	ch <- 1
}

func recive(ch chan int) {
	val := <-ch
	fmt.Println("In Recive")
	fmt.Println(val)
}

func ChannelSum() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	go sum(ch)
	time.Sleep(time.Second * 1)
}

func sum(ch <-chan int) {
	s := 0
	for val := range ch {
		s += val
	}
	fmt.Println(s)
}
