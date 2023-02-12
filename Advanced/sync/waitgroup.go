package main

import (
	"fmt"
	"sync"
)

func WaitGroup() {
	var wg sync.WaitGroup

	// Adding goroutine

	wg.Add(2)
	// Calling the goroutine
	go func() {
		defer wg.Done()
		fmt.Println("First GoRoutine is called")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Second GoRoutine is called")
	}()

	wg.Wait()

	fmt.Println("All the Waitgroup is completed")
}
