package concurrency

import (
	"fmt"
	"sync"
)

type PubSub struct {
	mu         sync.RWMutex
	subscriber map[string][]chan int
}

func NewPubSubModel() *PubSub {
	return &PubSub{
		subscriber: make(map[string][]chan int),
	}
}

func (ps *PubSub) SubScribe(topic string) <-chan int {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan int, 1)
	ps.subscriber[topic] = append(ps.subscriber[topic], ch)
	return ch
}

func (ps *PubSub) Publish(topic string, val int) {

	ps.mu.Lock()
	defer ps.mu.Unlock()

	for _, subscriber := range ps.subscriber[topic] {
		subscriber <- val
	}
}

func (ps *PubSub) Close(topic string, subCh <-chan int) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	_, found := ps.subscriber[topic]

	if !found {
		return
	}

	for i, subscriber := range ps.subscriber[topic] {

		if subscriber == subCh {
			close(subscriber)
			// removing this subscriber from the array
			ps.subscriber[topic] = append(ps.subscriber[topic][:i], ps.subscriber[topic][i+1:]...)
			break
		}
	}
}

func CallPubSubModel() {
	ps := NewPubSubModel()
	subscriber := ps.SubScribe("topic1")
	subscriber1 := ps.SubScribe("topic1")
	subscriber2 := ps.SubScribe("topic1")
	subscriber3 := ps.SubScribe("topic1")
	go func() {
		ps.Publish("topic1", 1)
	}()
	val := <-subscriber
	val1 := <-subscriber1
	val2 := <-subscriber2
	val3 := <-subscriber3

	fmt.Println(val)
	fmt.Println(val1)
	fmt.Println(val2)
	fmt.Println(val3)

	ps.Close("topic1", subscriber1)
	ps.Close("topic1", subscriber2)
	ps.Close("topic1", subscriber3)

}
