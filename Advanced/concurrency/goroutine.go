package concurrency

import (
	"fmt"
	"time"
)

func GoRoutineFunc(from string) {
	f("direct")
	go f("goroutine")
	go func(from string) {
		fmt.Println(from)
	}("going")

	time.Sleep(2 * time.Second)

	fmt.Println("Stop")

}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, " : ", i)
	}
}
