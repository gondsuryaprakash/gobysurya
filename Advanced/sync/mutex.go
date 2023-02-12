package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

func MutexWorkFlow() {

	var wg sync.WaitGroup
	// Adding waitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			mutex.Lock()
			count++
			mutex.Unlock()
		}
		fmt.Println("First GoRoutine is Done :", count)
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			mutex.Lock()
			count++
			mutex.Unlock()
		}
		fmt.Println("Secod GoRoutine is Done :", count)
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			mutex.Lock()
			count++
			mutex.Unlock()
		}
		fmt.Println("Third GoRoutine is Done :", count)
	}()

	wg.Wait()

	fmt.Println("All Goroutines are completed !!")
}
