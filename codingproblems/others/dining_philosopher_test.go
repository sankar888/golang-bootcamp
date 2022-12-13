package codingproblems

import (
	"testing"
	"time"
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
	
	go p1.eat()
	go p2.eat()
	go p3.eat()
	go p4.eat()
	go p5.eat()

	

}
