package main

import (
	"sync"
)

func main() {
	var mu sync.Locker = new(I)
	defer LockUnlock(mu)()
	println("doing")
}

// LockUnlock lock unlock
func LockUnlock(mu sync.Locker) (unlock func()) {
	mu.Lock()
	return mu.Unlock
}

// I i struct
type I struct{}

// Lock lock
func (i *I) Lock() {
	println("lock")
}

// Unlock unlock
func (i *I) Unlock() {
	println("unlock")
}
