package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Square(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i * i
	}

	close(c)
}

func CallSquare() {

	c := make(chan int)

	go Square(c)
	// use of for range

	for val := range c {
		fmt.Println("Channel value :", val)
	}
	// for {
	// 	val, ok := <-c
	// 	if ok == false {
	// 		fmt.Println("Channel stop sending message ", val, ok)
	// 		break
	// 	} else {
	// 		fmt.Println("Value of Channel :", val, ok)
	// 	}

	// }
}

// Buffered Channel

var start time.Time

func init() {
	start = time.Now()
}

func Squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num)
	}
}

func OddPrint(c chan int, wg *sync.WaitGroup) {
	for i := 0; i <= 10; i++ {
		odd := <-c
		fmt.Println("Odd Numbers :", odd)
		c <- odd
	}
	wg.Done()
}
func EvenPrint(c chan int, wg *sync.WaitGroup) {
	for i := 0; i <= 10; i++ {
		even := <-c
		fmt.Println("Even Numbers :", even)
		c <- even
	}
	wg.Done()
}

func PrintNumbers(c chan int) {
	select {
	case <-c:
		fmt.Println("Numbers ", <-c)

	}
}

func Print(c chan int) {
	for i := 0; i <= 20; i++ {
		number := <-c
		fmt.Println("Numbers Are : ", number)
	}
}
func CallBufferedChannel() {
	// c := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(2)
	even := make(chan int, 10)
	odd := make(chan int, 10)

	go OddPrint(odd, &wg)
	go EvenPrint(even, &wg)
	// go PrintNumbers()
	// go Print(c)

	for i := 0; i <= 20; i++ {

		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}

	}

	select {
	case res := <-even:
		fmt.Println("Even Res :", res)
	case res2 := <-odd:
		fmt.Println("Odd Res :", res2)

	}

	wg.Wait()

}

func Worker(i int, ch chan int) {
	fmt.Println("Started Working ", i)
	time.Sleep(time.Second * 2)
	fmt.Println("End task ", i)


	fmt.Printf("Cap %v , %v", cap(ch), i)


}

func CallWorker() {
	bufferedChannel := make(chan int, 3)
	for i := 0; i < 5; i++ {
		go Worker(i, bufferedChannel)
		bufferedChannel <- i
	}
	time.Sleep(time.Second * 7)
}
