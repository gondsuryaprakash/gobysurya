package channel

import (
	"fmt"
	"strings"
)

func PingPong(ping chan string, pong chan string) {

	for {
		message := <-ping
		pong <- fmt.Sprintf("%v!!", strings.ToUpper(message))
	}
}
func CallPingPong() {
	ping := make(chan string)
	pong := make(chan string)
	go PingPong(ping, pong)

	for {
		fmt.Print("->")
		// reciveing message from User
		var message string
		_, _ = fmt.Scanln(&message)
		if strings.ToLower(message) == "q" {
			break
		}
		ping <- message
		response := <-pong
		fmt.Println("Response : ", response)
	}

	fmt.Println("All Done closing the channel : ")
	close(ping)
	close(pong)
}
