package channel

import (
	"fmt"
	"time"
)

func Tickers() {
	tickers := time.NewTicker(time.Second * time.Duration(1))

	done := make(chan bool)
	go func() {

		for {
			select {
			case <-tickers.C:
				Print(tickers)
			case <-done:
				return
			}
		}
	}()
	time.Sleep(time.Second * 15)
	tickers.Stop()
	done <- false
	fmt.Println("Done")
}

func Print(t *time.Ticker) {
	fmt.Println("Surya is Don ! ", t)
}
