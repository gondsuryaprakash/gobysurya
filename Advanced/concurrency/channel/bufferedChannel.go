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

	// select {
	// case message1 := <-bufferedChannel:
	// 	fmt.Println(message1)
	// case message2 := <-bufferedChannel:
	// 	fmt.Println(message2)
	// }

}
