package channel

import (
	"fmt"
	"sync"
	"time"
)

func CreateUnBufferedChannel() {

	funcName := "CreateUnBufferedChannel"

	fmt.Println("Inside ", funcName)

	c := make(chan string)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		c <- "Hi I am from Channel"
		defer wg.Done()

	}()
	go func() {
		time.Sleep(2 * time.Second)
		// Communication will be only successfull when we have both receiver and emitter

		fmt.Println("Message form the Channel <- ", <-c)
		defer wg.Done()
	}()
	wg.Wait()
}
