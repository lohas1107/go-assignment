package main

import "sync"

type User struct {
	ID      uint64
	Balance uint64
	Lock    sync.Mutex
}

func transfer(from *User, to *User, amount uint64) {
	from.Lock.Lock()
	to.Lock.Lock()
	defer from.Lock.Unlock()
	defer to.Lock.Unlock()

	if from.Balance >= amount {
		from.Balance -= amount
		to.Balance += amount
	}
}

func main() {
	userA := User{
		ID: 1, Balance: 10e10,
	}
	userB := User{
		ID: 2, Balance: 10e10,
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _ = range [10e10]uint64{} {
			transfer(&userA, &userB, 1)
		}
	}()

	go func() {
		defer wg.Done()
		for _ = range [10e10]uint64{} {
			transfer(&userB, &userA, 1)
		}
	}()
	wg.Wait()
}
