package channel

import "fmt"

func SumOfSquare() {
	number := make(chan int)

	quit := make(chan int)
	sum := 0
	go func() {
		for i := 0; i < 10; i++ {
			sum += <-number
		}
		quit <- 0
	}()

	SquareOfNumber(quit, number)
	fmt.Println("Sum Of Number :", sum)
}

func SquareOfNumber(quit, c chan int) {

	y := 1
	for {
		select {
		case c <- (y * y):
			y++
		case <-quit:
			return
		}
	}
}
