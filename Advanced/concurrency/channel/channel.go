package channel

import "fmt"

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



