package codingproblems

import (
	"testing"
	"time"
	"context"
	"sync"
)

func TestDiningPhilosophers(t *testing.T) {
	var c1 chopstick = NewChopstick()
	var c2 chopstick = NewChopstick()
	var c3 chopstick = NewChopstick()
	var c4 chopstick = NewChopstick()
	var c5 chopstick = NewChopstick()

	var p1 *philosopher = NewPhilosopher("p1", c1, c2)
	var p2 *philosopher = NewPhilosopher("p2", c2, c3)
	var p3 *philosopher = NewPhilosopher("p3", c3, c4)
	var p4 *philosopher = NewPhilosopher("p4", c4, c5)
	var p5 *philosopher = NewPhilosopher("p5", c5, c1)
	
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)
	var wg sync.WaitGroup
	wg.Add(5)
	go p1.eat(ctx, &wg)
	go p2.eat(ctx, &wg)
	go p3.eat(ctx, &wg)
	go p4.eat(ctx, &wg)
	go p5.eat(ctx, &wg)
	wg.Wait()

	t.Logf("p1 : %d, p2 : %d, p3 : %d, p4 : %d, p5 : %d\n", p1.ate(), p2.ate(), p3.ate(), p4.ate(), p5.ate())
}
