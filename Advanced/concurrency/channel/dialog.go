package channel

import (
	"fmt"
	"os"
	"time"
)

// Two channel can dialog through channel

type Ball uint64

func Play(playerName string, table chan Ball) {
	var lastValue Ball = 1

	for {
		number := <-table
		fmt.Println(playerName, number)
		number += 1
		if number < lastValue {
			os.Exit(0)
		}
		lastValue = number
		table <- number
		time.Sleep(time.Second)
	}
}

func CallPlay() {
	ch := make(chan Ball)

	go func() {
		ch <- 1

	}()
	go Play("A", ch)
	Play("B", ch)
}

func Printer(funcName string, ch chan int) {

}
