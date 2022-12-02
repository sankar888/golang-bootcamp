package main

import (
	"testing"
)

func main() {
	fmt.Println("main function Exited")
}

// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
// Data flows in the direction of the operatoe <-
// By default, sends and receives block until the other side is ready.
// This allows goroutines to synchronize without explicit locks or condition variables.
func TestChannelsBasicUsage(t *testing.T) {
	//create a channel
	var c chan int
	t.Logf("type of channel %T, value of channel %v\n", c, c)
	c = make(chan int, 2) //bufferred channel

	c <- 100
	defer close(c)

	//read from channel
	var i int = <- c
	t.Logf("value read from channel %d \n", i)
}