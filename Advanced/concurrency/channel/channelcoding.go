package channel

import "fmt"

// Creating function which take channel as an argument and send result through channel
func MultiPlyByTwo(n int, out chan<- int) {
	product := n * 2
	out <- product
}

func CallChannel() {
	out := make(chan int)

	// Calling goroutine functions and sending channel
	go MultiPlyByTwo(10, out)

	// Recieve message through channel
	fmt.Println(<-out)
}

func CallDirectionalChannel() {

	out1 := make(chan int)
	in1 := make(chan int)

	go MultiplyChannel(in1, out1)
	go MultiplyChannel(in1, out1)
	go MultiplyChannel(in1, out1)

	go func() {
		// Sending Data to channel
		in1 <- 1
		in1 <- 2
		in1 <- 3
		in1 <- 4
	}()

	// Recieving data from channel
	fmt.Println(<-out1)
	fmt.Println(<-out1)
	fmt.Println(<-out1)
	fmt.Println(<-out1)

}

func MultiplyChannel(in <-chan int, out chan<- int) {

	// Getting data through channel
	num := <-in

	product := num * 2

	// Sending data to channel
	out <- product
}
