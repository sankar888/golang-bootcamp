package learncontext

import (
	"testing"
	"context"
	"time"
)

//https://blog.golang.org/context
func TestContextBasics(t *testing.T) {
	//how to create a context
	var ctx context.Context = context.Background()
	t.Logf("context type: %T, value: %v, pointer: %p \n", ctx, ctx, ctx)
	callMe(ctx, t) //to test if ctx is a pointer or reference type. ctx is a pointer

	//what will calls to empty context do ?
	deadline, ok := ctx.Deadline() //deadline returns (deadline time.Time, ok bool). Deadline returns ok==false when no deadline is set
	t.Logf("calling Deadline on empty context, deadline: %v, exceeded: %v \n", deadline, ok)

	ch := ctx.Done() //since done is called on empty context , it returns nil
	// Done returns a channel that's closed when work done on behalf of this
	// context should be canceled. Done may return nil if this context can
	// never be canceled. Successive calls to Done return the same value.
	// The close of the Done channel may happen asynchronously,
	// after the cancel function returns.
	//
	// WithCancel arranges for Done to be closed when cancel is called;
	// WithDeadline arranges for Done to be closed when the deadline
	// expires; WithTimeout arranges for Done to be closed when the timeout
	// elapses.
	//
	// Done is provided for use in select statements:
	//
	//  // Stream generates values with DoSomething and sends them to out
	//  // until DoSomething returns an error or ctx.Done is closed.
	//  func Stream(ctx context.Context, out chan<- Value) error {
	//      for {
	//              v, err := DoSomething(ctx)
	//              if err != nil {
	//                      return err
	//              }
	//              select {
	//              case <-ctx.Done():
	//                      return ctx.Err()
	//              case out <- v:
	//              }
	//      }
	//  }
	//
	// See https://blog.golang.org/pipelines for more examples of how to use
	// a Done channel for cancellation.
	t.Logf("calling Done() on empty context, should return nil, ctx: %v\n", ch)

	err := ctx.Err()
	// If Done is not yet closed, Err returns nil.
	// If Done is closed, Err returns a non-nil error explaining why:
	// Canceled if the context was canceled
	// or DeadlineExceeded if the context's deadline passed.
	// After Err returns a non-nil error, successive calls to Err return the same error.
	t.Logf("calling Err() on empty context, should return nil, err: %v\n", err)

	val := ctx.Value("somekey")
	// Value returns the value associated with this context for key, or nil
	// if no value is associated with key. Successive calls to Value with
	// the same key returns the same result.
	t.Logf("calling Value(key) n background context, should return nil val: %v\n", val)
}

func callMe(ctx context.Context, t *testing.T) {
	t.Logf("is context reference type : context type: %T, value: %v, pointer: %p \n", ctx, ctx, ctx)
}


// this test demonstrates how to create derived context
func TestCreateDerivedContext(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc() //if the program executes normally this will run but will have no effect
	//wait for 10 sec and then cancel
	go func() {
		timer := time.NewTimer(time.Second * 10)
		<-timer.C 
		cancelFunc()
	}()

	// when cancelFunc is called,
	// ctx.Done() channel will be closed and will not block
	// ctx.Err() will give some non-nil error
	// ctx.Deadline() will return (start of time, false), since it is not a DealineContext
	err := longRunningTask(ctx, t)
	if err != nil {
		t.Log("long running task ended", err)
	}

	//auto timeout after 5 sec
	ctx, cancelFunc = context.WithTimeout(context.Background(), (time.Second * 5))
	defer cancelFunc() //if the program executes normally this will run but will have no effect
	err = longRunningTask(ctx, t)
	if err != nil {
		t.Log("long running task timedout", err)
	}

	//deadline context
	ctx, cancelFunc = context.WithDeadline(context.Background(), time.Now().Add(time.Second * 3))
	defer cancelFunc() //if the program executes normally this will run but will have no effect
	err = longRunningTask(ctx, t)
	if err != nil {
		t.Log("long running task deadline exceeded", err)
	}
}

// this function will print numbers untill cancelled by context
func longRunningTask(ctx context.Context, t *testing.T) error {
	t.Log("long running task.")
	var i int = 0
	for {
		select {
		case <- ctx.Done():
			t.Log("cancel context cancelled. long running task quitting.") 
			return ctx.Err()
		default:
			i = i+1
			t.Log("long running task running", i)
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

// the below test demonstrates the uage of select statement
func TestSelectUsage(t *testing.T) {
	//TODO: implement practical uses of select
}

