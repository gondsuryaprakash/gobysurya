package channel

import (
	"fmt"
)

func NewChannel() {

	bufferedChannel := make(chan string, 2)

	go func() {
		bufferedChannel <- "First Message"
		bufferedChannel <- "Second Message"
	}()

	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

}
