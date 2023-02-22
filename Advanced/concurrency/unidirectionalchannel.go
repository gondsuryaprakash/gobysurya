package concurrency

import "fmt"

/*
Two type of channel in Unidirectional
1. Recieve only Channel
2. Send only Channel

roc := make(<-chan int)
soc := make(chan<- int)

*/

func GreetROC(roc <-chan string) {
	fmt.Println("Hello From the :", <-roc)
}

func CallROC() {

	c := make(chan string)

	go GreetROC(c)

	c <- "Surya "

	fmt.Println("main() stopped")
}

func GreetSOC(soc chan<- string) {
	soc <- "Surya"
	fmt.Println("Sending data from GreetSOC")
}

func CallSOC() {

	c := make(chan string)

	go GreetROC(c)

	fmt.Println("Taking the value from the : ", <-c)

	fmt.Println("main() stopped")
}
