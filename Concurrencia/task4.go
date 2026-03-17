package main

import (
	"fmt"
	"sync"
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	id             int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
	hostRequest    chan chan bool // Channel to ask host for permission
	hostRelease    chan bool      // Channel to tell host we are done
	wg             *sync.WaitGroup
}

func (p Philosopher) eat() {
	defer p.wg.Done()

	for i := 0; i < 3; i++ {
		permissionChan := make(chan bool)
		p.hostRequest <- permissionChan
		<-permissionChan // Wait block until host grants permission

		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		// C. Start eating
		fmt.Printf("starting to eat %d\n", p.id)

		fmt.Printf("finishing eating %d\n", p.id)

		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()

		p.hostRelease <- true
	}
}

func host(request <-chan chan bool, release <-chan bool) {
	eating := 0
	for {
		if eating < 2 {
			select {
			case respChan := <-request:
				eating++
				respChan <- true // Grant permission
			case <-release:
				eating--
			}
		} else {
			<-release
			eating--
		}
	}
}

func main() {
	Csticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		Csticks[i] = new(Chopstick)
	}

	hostReq := make(chan chan bool)
	hostRel := make(chan bool)

	go host(hostReq, hostRel)

	var wg sync.WaitGroup
	wg.Add(5)

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1, // Numbered 1 to 5
			leftChopstick:  Csticks[i],
			rightChopstick: Csticks[(i+1)%5], // The right chopstick is the next one in the circle
			hostRequest:    hostReq,
			hostRelease:    hostRel,
			wg:             &wg,
		}

		go philosophers[i].eat()
	}

	wg.Wait()
}
