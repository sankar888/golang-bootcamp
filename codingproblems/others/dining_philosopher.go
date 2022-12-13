package codingproblems

import (
	"fmt"
	"time"
	"context"
	"sync"
)

type chopstick chan int

func NewChopstick() chopstick {
	c := make(chan int, 1)
	c <- 1
	return c
}

type philosopher struct {
	name  string
	left  chopstick
	right chopstick
	food   int
}

func NewPhilosopher(name string, left chopstick, right chopstick) *philosopher {
	return &philosopher {
		name:  name,
		left:  left,
		right: right,
	}
}

func (p *philosopher) eat(ctx context.Context, wg *sync.WaitGroup) {
	loop:
	for {
		select {
		case <-p.left:
			//wait for some time for right to be available
			select {
			case <- p.right:
				for i := 0; i < 10; i++ {
					p.food += 1
					fmt.Printf("%s acquired both left and right. Eating.. %d\n", p.name, p.food)
				}
				p.right <- 1
				p.left <- 1
			case <- time.After(time.Millisecond * 200):
				fmt.Printf("%s acquired left waited 200ms for right to be available. trying again.\n", p.name)
				p.left <- 1
			}

		case <-p.right:
			//wait for sometime to left to be available
			select {
			case <- p.left:
				for i := 0; i < 10; i++ {
					p.food += 1
					fmt.Printf("%s acquired both left and right. Eating.. %d\n", p.name, p.food)
				}
				p.right <- 1
				p.left <- 1
			case <- time.After(time.Millisecond * 200):
				fmt.Printf("%s acquired right waited 200ms for left to be available. trying again.\n", p.name)
				p.right <- 1
			}

		case <-time.After(time.Second): //wait for sometime either right or left to be available
			fmt.Printf("%s waited 1s for left or right to be available. trying again.\n", p.name)

		case <-ctx.Done():
			fmt.Printf("%s Done eating. exiting.\n", p.name)
			wg.Done()
			break loop
		}
	}
}

func (p *philosopher) ate() int {
	return p.food
}