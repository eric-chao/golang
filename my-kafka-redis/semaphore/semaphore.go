package semaphore

import "sync"

// Semaphore implements the signal mechanism management
type Semaphore struct {
	Threads chan int
	Wg      sync.WaitGroup
}

// NewSemaphore returns a new semaphore
func NewSemaphore(n int) *Semaphore {
	inst := new(Semaphore)
	inst.Threads = make(chan int, n)
	// inst.Wg=sync.WaitGroup{}
	return inst
}

// P is a primitive operation that requests the allocation of a unit resource
func (sem *Semaphore) Acquire() {
	sem.Threads <- 1
	sem.Wg.Add(1)
}

// V is a primitive operation that releases a unit of resources
func (sem *Semaphore) Release() {
	sem.Wg.Done()
	<-sem.Threads
}

// Wait waits for all unit resources to be released
func (sem *Semaphore) Wait() {
	sem.Wg.Wait()
}
