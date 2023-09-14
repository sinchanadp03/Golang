package main

import (
	"fmt"
	"sync"
)

type Philosopher struct {
	id                  int
	leftFork, rightFork *sync.Mutex
}

func (p Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.leftFork.Lock()
		p.rightFork.Lock()

		fmt.Printf("Philosopher %d is eating\n", p.id)

		p.leftFork.Unlock()
		p.rightFork.Unlock()

		fmt.Printf("Philosopher %d is thinking\n", p.id)
	}
}

func main() {
	var num int
	fmt.Print("Enter:")
	fmt.Scan(&num)
	var wg sync.WaitGroup
	forks := make([]*sync.Mutex, num)
	philosophers := make([]Philosopher, num)

	for i := 0; i < num; i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < num; i++ {
		philosophers[i] = Philosopher{
			id:        i + 1,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%num],
		}
	}

	for i := 0; i < num; i++ {
		wg.Add(1)
		go philosophers[i].dine(&wg)
	}

	wg.Wait()
}
