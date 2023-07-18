package channel

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func OEPrinter() {
	var wg sync.WaitGroup
	wg.Add(2)

	number := make(chan int)
	done := make(chan bool)

	go PrintEvens(number, done, &wg)
	go PrintOdds(number, done, &wg)

	for i := 0; i <= 10; i++ {
		number <- i
		<-done
	}
	close(number)
	close(done)
	wg.Wait()

}

func PrintEvens(number chan int, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range number {
		if num%2 == 0 {
			fmt.Println("Printing from the Evens", num)
			done <- true
		} else {
			number <- num
		}
	}

}

func PrintOdds(number chan int, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range number {
		if num%2 != 0 {
			fmt.Println("Printing from the Odds", num)
			done <- true
		} else {
			number <- num
		}
	}

}

// Print Ping Pong Using one counter
func PrintPingPong() {
	// var wg sync.WaitGroup
	// wg.Add(2)
	pinger := make(chan int)
	ponger := make(chan int)

	go PrintPing(pinger, ponger)
	go PrintPong(pinger, ponger)

	pinger <- 1

	time.Sleep(time.Second * 3)

	os.Exit(0)

}

func PrintPing(pinger <-chan int, ponger chan<- int) {

	for {
		val := <-pinger
		fmt.Println("Ping :)", val)
		time.Sleep(time.Millisecond * 200)
		ponger <- val + 1
	}

}

func PrintPong(pinger chan<- int, ponger <-chan int) {
	for {
		val := <-ponger
		fmt.Println("Pong :)", val)
		time.Sleep(time.Millisecond * 500)
		pinger <- val + 1
	}

}

// Print Odd even Using one counter

func PrintOddEvenUsingCounter() {

	even := make(chan int)
	odd := make(chan int)

	go PrintEvenWithCounter(even, odd)
	go PrintOddWithCounter(even, odd)

	even <- 0
	time.Sleep(5 * time.Second)
	os.Exit(0)
}

func PrintEvenWithCounter(even <-chan int, odd chan<- int) {
	for {
		val := <-even
		fmt.Println("Even ", val)
		time.Sleep(time.Millisecond * 200)
		odd <- val + 1
	}
}

func PrintOddWithCounter(even chan<- int, odd <-chan int) {

	for {
		val := <-odd
		fmt.Println("Odd ", val)
		time.Sleep(time.Millisecond * 200)
		even <- val + 1
	}
}
