package main

import (
	"time"
	"sync"
	"testing"
)

func TestWaitGroupBasic(t *testing.T) {
	t.Log("main program starts.")
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(1) //wait untill this many Done is called
	//spin up a go routine
	go func(){
		for i := 0; i < 10; i++ {
			t.Log("go routine running, out ->", i)
			time.Sleep(time.Duration(2 * time.Second))
		}
		t.Log("go routine completed.")
		wg.Done()
	}()

	t.Log("main program waiting for the go routine to complete.")
	wg.Wait() //will bloack and wait for wg count to zero
	t.Log("main progeram completed.")
}


// This test will forever because the waitgroup counter is set to 2 and only one go routine calls Done()
// so the waithroup will wait until the counter is zero. thiat is forerver.
// This means, no matter the state of the routine Done should be called even in error
func TestWGPanic(t *testing.T) {
	t.Log("main program starts.")
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(1) //wait untill this many Done is called
	
	//spin up a go routine
	go func(){
		defer func(){
		if r := recover(); r != nil {
			t.Log("recovered from panic:", r)
		}
		wg.Done()
		}()
		for i := 0; i < 2; i++ {
			t.Log("go routine running, out ->", i)
			time.Sleep(time.Duration(2 * time.Second))
			panic("go routine panics")
		}
		t.Log("go routine completed.")
	}()

	t.Log("main program waiting for the go routine to complete.")
	wg.Wait() //will bloack and wait for wg count to zero
	t.Log("main program completed.")
}

// Does defer function executes even in case of panic. Yes
func TestIfPanicExecutesDefer(t *testing.T) {
	t.Log("programs starts")
	defer t.Log("defered clean up scripts executed.")
	panic("program panics")
}


// does panic bubble up. Yes
func TestIfPanicBubbleUp(t *testing.T) {
	defer func(){
		if r := recover(); r != nil {
			t.Log("recovered from panic:", r)
		}
	}()
	t.Log("main program")
	t.Log("calling callMe()")
	callMe()
	t.Log("this statement should not execute.")
}

func callMe() {
	panic("callMe panics")
}


// test if the waitgroup is a reference type like slice, channel
// no waitgroup is not reference type
func TestIfWGIsReferencerType(t *testing.T) {
	var wg sync.WaitGroup = sync.WaitGroup{}
	t.Logf("main function. wait group type: %T, pointer: %p value: %v \n", wg, &wg, wg)
	t.Log("wait group passed to another function.")
	callMeWG(wg, t)
}

func callMeWG(wg sync.WaitGroup, t *testing.T) {
	t.Logf("wait group type: %T, pointer: %p value: %v \n", wg, &wg, wg)
}


// wait group whose counter is not set to zero by other go routines waits forever
// wg.Add() adds the number to already exiting counter
func TestWGWaitForever(t *testing.T) {
	t.Log("program starts")
	wg := sync.WaitGroup{}
	wg.Wait() //this will not wait since no counter is set
	
	t.Log("waitgroup with counter 1 added")
	wg.Add(1)

	t.Log("waitgroup with counter -1 added. so now wait won't block")
	wg.Add(-1)
	wg.Wait()
	t.Log("previous wait didn't block")

	wg.Add(1)
	wg.Add(1) //if this is commented the test will pass

	wg.Done()
	wg.Wait() //now it will wait forever, since there is no go routine to call Done()
	t.Log("program ends")
}