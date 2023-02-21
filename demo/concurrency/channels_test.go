package learnchannel

import (
	"testing"
	"time"
)

// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
// Data flows in the direction of the operator <-
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
	var i int = <-c
	t.Logf("value read from channel %d \n", i)
}

// channel send and receive will block untill the other side is ready
// for unbuferred channel, send or receive without another goroutine will result in deadlock
// what happens if the goroutine doesnot send eny event and main routine blocks for receive in an unbufferred channel
func TestChannelDeadlock(t *testing.T) {
	t.Log("main program starts.")
	var ch chan int = make(chan int)

	//go routine runs but not sendinging anything
	go func() {
		t.Log("side go routine running.")
	}()
	close(ch) //channel closed. so receive wont block
	//close(ch) //closing an already closed channel will cause panic

	count := <-ch //will not cause deadlock, but blocks to receive an int, if channel is not closed, comment the close and see how it behaves

	//send in a closed channel will panic
	//ch <- 20
	t.Log("main program ends. count:", count)
}

// The below test illustrates that chnnels are like queues (not like a broadcast).
// the received data is taken out of channels
// channels will not act like a brodcast write once read many.
func TestChannelsWithMultipleRoutines(t *testing.T) {
	t.Log("main program starts.")

	var signal chan int = make(chan int)
	for i := 0; i < 10; i++ {
		go func(num int) {
			t.Log("go routine", num, "started. but waiting for start signal")
			<-signal //will block until some one sends something
			t.Logf("go routine %d ends. \n", num)
		}(i)
	}

	time.Sleep(time.Second * 2) //wait for some time before starting all routines
	signal <- 200               //channels are like pipe, since we have sent only one 200 , only one go routine will pick it up and run
	t.Log("sig <- 200. starts all go routine")
	close(signal)               //closing the channel act like a broadcast, all routines will unblock
	time.Sleep(time.Second * 3) //wait for some time for all routines to complete
	t.Log("main program ends.")
}
