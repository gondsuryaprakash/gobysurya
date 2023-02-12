package channel

import "fmt"

func Square(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i * i
	}

	close(c)
}

func CallingChannel() {
	fmt.Println("CallingChannel started ! ")
	chanel := make(chan int)
	go Square(chanel)

	for {

		val, ok := <-chanel

		if ok == false {
			fmt.Println("Loop break")
			break
		} else {
			fmt.Println("Square of number: ", val, ok)
		}

	}

}
