package faninfanout

import (
	"fmt"
	"time"
)

func Generator(data string) <-chan string {

	channel := make(chan string)

	go func() {

		for {
			channel <- data
			time.Sleep(1000)
		}
	}()

	return channel
}

func FanIn() {
	c1 := Generator("Hello")
	c2 := Generator("There")

	fanIn := make(chan string)

	go func() {
		for {
			select {
			case str1 := <-c1:
				fanIn <- str1
			case str2 := <-c2:
				fanIn <- str2
			}
		}
	}()

	go func() {
		for {
			fmt.Println(<-fanIn)
		}

	}()
}

type Processor struct {
	JobChannel chan string
	done       chan *Worker
	workers    chan []*Worker
}

type Worker struct {
	name string
}

func (w *Worker) ProcessJob(data string, done chan *Worker) {
	go func() {
		fmt.Println("Working on Data ", data, w.name)

		done <- w
	}()
}

func GetProcessor() *Processor {
	return &Processor{}
}

func ChannelPractice() {
	c := make(chan int)
	// sending data to the channel
	go func() {
		c <- 1
		time.Sleep(1 * time.Second)
	}()
	data := <-c
	fmt.Println(data)

}
