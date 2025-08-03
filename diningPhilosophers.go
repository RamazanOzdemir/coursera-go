package main

import (
	"fmt"
	"sync"
	"time"
)

// Dining Philosophers problem is a classic synchronization problem
// where philosophers alternate between thinking and eating.
// The challenge is to avoid deadlock and ensure that all philosophers can eat.

type CopStick struct {
	sync.Mutex
}

type Philosopher struct {
	id int 
	left, right *CopStick
}

func (p *Philosopher) eat(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()

	// Wait for permission to eat
	<-ch

	
	// Pick up left and right chopsticks
	p.left.Lock()
	p.right.Lock()

	

	// Simulate eating
	fmt.Printf("starting to eat %d\n", p.id)
	time.Sleep(300 * time.Millisecond) // Simulate time taken to eat
	fmt.Printf("finishing eating %d\n", p.id)
	
	// Put down chopsticks
	p.right.Unlock()
	p.left.Unlock()

	// Allow another philosopher to eat
	ch <- true
}

type Host struct {
}

func (h *Host) initializeTale(p *[]Philosopher, sticks *[]CopStick) {
	for i := 0; i < 5; i++ {
		(*p)[i] = Philosopher{
			id: i,
			left: &(*sticks)[i],
			right: &(*sticks)[(i+1)%5],
		}
	}
}


func main() {
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	
	// permission channel to control when max 2 philosophers can eat  at same time 
	ch := make(chan bool, 2)
	ch <- true // Initially allow one philosopher to eat
	ch <- true // Allow a second philosopher to eat

	// Initialize the host, chopsticks, and philosophers
	h:= Host{}
	sticks := make([]CopStick, 5)
	philosophers := make([]Philosopher, 5)

	h.initializeTale(&philosophers, &sticks)


	// Start the philosophers eating
	// Each philosopher could eat 3 times
	// 5 philosophers, 2 at a time
	// 5 philosophers * 3 times each = 15 total eating
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go philosophers[i%5].eat(&wg, ch)
	}
	wg.Wait()
	// fmt.Println("All philosophers have finished eating.")

}