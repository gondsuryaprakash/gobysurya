package mutex

import (
	"fmt"
	"sync"
)

type Account struct {
	Balance int
	lock    sync.Mutex
}

func (c *Account) GetBalance() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.Balance
}

func (c *Account) Withdraw(v int) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if v > c.Balance {
		fmt.Println("Insufficent Balance")
	} else {
		c.Balance -= v
	}

}

