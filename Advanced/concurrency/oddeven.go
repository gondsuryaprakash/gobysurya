package concurrency

import "fmt"

func CallPrintOddEven() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go Send(even, odd, quit)
	Recieve(even, odd, quit)

	fmt.Println("main() stopped")
}

func Send(even, odd chan<- int, quit chan<- bool) {

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	quit <- true
}

func Recieve(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case ev := <-even:
			fmt.Println("Even Number :", ev)
		case od := <-odd:
			fmt.Println("Odd Number :", od)

		case qu := <-quit:
			fmt.Println("Quit the GoRoutine :", qu)
			return
		}
	}
}
