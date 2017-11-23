package bank3

import "sync"

var (
	//mu sync.Mutex // guards balance
	mu sync.RWMutex // multiple readers, single writer lock
	balance int
)

func Deposit(amount int) {
	mu.Lock() // write lock
	defer mu.Unlock()
	balance = balance + amount
	//mu.Unlock()
}
func Balance() int  {
	mu.RLock() // read lock
	defer mu.RUnlock()
	//mu.Lock()
	b := balance
	//mu.Unlock()
	return b
}
