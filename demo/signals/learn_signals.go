package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
)

// A main program which runs and blocks
// This program demonstrates the use of go signals
func main() {
	// go receives incoming singla in a channel
	// the channel should be bufferred to receive the no of signals to be handled
	fmt.Println("main program starts.")
	var ch chan os.Signal = make(chan os.Signal)
	signal.Notify(ch) //Nofify all signals


	go func() {
		fmt.Println("inside go routine. waiting for 2 sec.")
		time.Sleep(time.Second * 2)
		signal.Reset() //reset all signals. No furhther signals will be caught
		//signal.Stop(ch) //this will clear 
		fmt.Println("stopped the signals channel. this should unblock any wait.")
	}()

	fmt.Println("waiting for any signal")
	//the below code will block and wait, but a receive in unbufferred channel with no other goroutine should case deadlock error
	//where is the other goroutine running ?
	//signals.Notify might look for signals in a separate go routine
	sig := <-ch 
	fmt.Println("main program ends received signal", sig)
}