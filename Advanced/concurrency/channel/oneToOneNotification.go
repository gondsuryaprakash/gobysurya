package channel

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
	"time"
)

func OneToOneNotificationBySending() {
	values := make([]byte, 32*1024*1024)

	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{})

	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})

		done <- struct{}{}
	}()

	<-done
	fmt.Println("values :", values)
	fmt.Println(values[0], values[len(values)-1])
}

func OneToOneNotificationByRecieving() {

	done := make(chan struct{})

	go func() {
		fmt.Println("Hello")

		time.Sleep(time.Second * 2)
		<-done
	}()

	done <- struct{}{}
	fmt.Println("World")
}
