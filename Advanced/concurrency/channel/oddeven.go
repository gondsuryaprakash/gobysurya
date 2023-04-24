package channel

import (
	"fmt"
	"sync"
)

func PrintOdd(number chan int, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range number {
		if n%2 != 0 {
			fmt.Println("Odd :", n)
			done <- true
		} else {
			number <- n
		}

	}

}

func PrintEven(number chan int, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range number {
		if n%2 == 0 {
			fmt.Println("Even :", n)
			done <- true
		} else {
			number <- n
		}
	}
}

func OddEvenPrinter() {
	var wg sync.WaitGroup

	wg.Add(2)
	number := make(chan int)
	isDone := make(chan bool)

	go PrintEven(number, isDone, &wg)
	go PrintOdd(number, isDone, &wg)

	for i := 0; i <= 10; i++ {
		number <- i
		<-isDone
	}

	close(number)
	close(isDone)

	wg.Wait()

}
