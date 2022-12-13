package learnatomic

import (
	"testing"
	"sync"
	"sync/atomic"
	"github.com/sankar888/golang-bootcamp/demo/common"
)

/**
 * Refer: https://pkg.go.dev/sync/atomic@go1.19.4 
 * 
 */

// this test demonstrates the need for synchronization
func TestWithoutSynchronization(t *testing.T) {
	common.Start("counter without synchronization")
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	counter, routines  := 0, 100
	for i := 0; i < routines; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			counter += 1
		}()
	}
	wg.Wait()
	if counter != routines {
		t.Logf("expected: %d , got: %d\n", routines, counter)
		t.Fail()
	}
	common.End()
}

func TestAtomicBasics(t *testing.T) {
	common.Start("Atomic Add")
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	var counter int32 = 0
	routines := 100
	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}
	wg.Wait()
	if int(counter) != routines {
		t.Logf("testcase failed. expected: %d , got: %d\n", routines, counter)
		t.Fail()
	} else {
		t.Logf("testcase success. expected: %d , got: %d\n", routines, counter)
	}
	common.End()

	common.Start("It is not recommended to use atomic package directly. Use other sync types if possible")
	common.End()	
}

