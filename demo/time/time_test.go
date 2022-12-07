package learntime

import (
	"testing"
	"time"
	"context"
)

// Introduces the basic concepts of time.Time
func TestTimeBasic(t * testing.T) {
	// A Time represents an instant in time with nanosecond precision.
    // Programs using times should typically store and pass them as values, not pointers.
    
    // An time instance can be created using the following methods
    // func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	// func Now() Time
	// func Parse(layout, value string) (Time, error)
	// func ParseInLocation(layout, value string, loc *Location) (Time, error)
	// func Unix(sec int64, nsec int64) Time
	// func UnixMicro(usec int64) Time
	// func UnixMilli(msec int64) Time
    var now time.Time = time.Now()
    t.Logf("reference check: type: %T, value: %v, pointer %p \n", now, now, &now)
    referenceCheck(now, t) //time is a value type not a pointer

    //explore various ways to create time
    //bday := time.Date(1995, time.June, 23) // this will fail all are mandatory arguments
    bday := time.Date(1995, time.June, 23, 0, 0, 0, 0, time.Local)
    year, month, day := bday.Date()
    t.Logf("my birthday is, all value: %v date: %d-%d-%d \n", bday, year, month, day)

    //from unix timestamp
    var ts int64 = 1670133020000
    now = time.UnixMilli(ts)
    t.Logf("time from ts %d is %v. time instant to unix %d \n", ts, now, now.UnixMilli())

    ts = 0 //should be Jan 1, 1970
    now = time.UnixMilli(ts)
	t.Logf("time from ts %d is %v. time instant to unix %d \n", ts, now, now.UnixMilli())

	//the unix epoch is zone independent.
	ts = 1670133020000
	now = time.UnixMilli(ts)
	utc := now.UTC()
	t.Logf("time %v from ts %d, same in UTC time %v, same UTC in ts %d \n", now, ts, utc, utc.UnixMilli())

	//create an instance of time from string representation, the below functions does that
	// func Parse(layout, value string) (Time, error)
	// func ParseInLocation(layout, value string, loc *Location) (Time, error)
	//Year: "2006" "06"
	//Month: "Jan" "January" "01" "1"
	//Day of the week: "Mon" "Monday"
	//Day of the month: "2" "_2" "02"
	//Day of the year: "__2" "002"
	//Hour: "15" "3" "03" (PM or AM)
	//Minute: "4" "04"
	//Second: "5" "05"
	//AM/PM mark: "PM"
	
	//Numeric time zone offsets format as follows:

	//"-0700"     ±hhmm
	//"-07:00"    ±hh:mm
	//"-07"       ±hh
	//"-070000"   ±hhmmss
	//"-07:00:00" ±hh:mm:ss
	tcases := []struct {
		layout string
		time   string
	}{ 
		{
			layout: "2006-01-02",
		    time:   "2022-12-05",
		},
		{
			layout: time.RFC3339,
			time:   "2022-12-05T06:06:00+05:30",
		},
	}

	for _, tc := range tcases {
		now, err := time.Parse(tc.layout, tc.time)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Parsing:str %s parsed to time %v with layout %s, adding 1 hr to it %v \n", tc.time, now, tc.layout, now.Add(time.Duration(time.Hour)))	
	}

	for _, tc := range tcases {
		now, err := time.ParseInLocation(tc.layout, tc.time, time.Local)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Parsing:str %s parsed to time %v with layout %s, adding 1 hr to it %v \n", tc.time, now, tc.layout, now.Add(time.Duration(time.Hour)))	
	}

	//time arithmetic operation Add(), Sub(), Before(), After() 
	//the difference between 5 hrs from now and 4 hrs from now is 1  hr
	now = time.Now()
	plus5 := now.Add(time.Hour * 5)
	plus4 := now.Add(time.Hour * 4)
	oneHR := plus5.Sub(plus4) //gives a duration
	if time.Hour != oneHR {
		t.Logf("test case failed, expeced : %v, got %v \n", time.Hour, oneHR)
		t.Fail()
	} else {
		t.Logf("test case passed, expeced : %v, got %v \n", time.Hour, oneHR)
	}

	if got := plus5.After(plus4); !got {
		t.Logf("test case failed, expeced : %v, got %v \n", true, got)
		t.Fail()
	} else {
		t.Logf("test case passed, expeced : %v, got %v \n", true, got)
	}

	if got := plus4.Before(plus5); !got {
		t.Logf("test case failed, expeced : %v, got %v \n", true, got)
		t.Fail()
	} else {
		t.Logf("test case passed, expeced : %v, got %v \n", true, got)
	}

	if got := plus4.Before(plus4); got {
		t.Logf("test case failed, expeced : %v, got %v \n", false, got)
		t.Fail()
	} else {
		t.Logf("test case passed, expeced : %v, got %v \n", false, got)
	}

	//Locale info from time
	now = time.Now() //picks up default timezone +05:30
	if name, offset := now.Zone(); name != "IST" || float64(offset) != time.Duration(time.Minute * 30  + time.Hour * 5).Seconds() {
		t.Logf("test case failed!, expected zone and offset : %s %d, got %s %d \n", "IST", 19800, name, offset)
		t.Fail()
	} else {
		t.Logf("test case passed., expected zone and offset : %s %d, got %s %d \n", "IST", 19800, name, offset)
	}
}

func referenceCheck(now time.Time, t *testing.T) {
	t.Logf("reference check: type: %T, value: %v, pointer %p \n", now, now, &now)
}


// this function will demonstrates the basic usage of time.Timer
// The Timer type represents a single event. When the Timer expires, the
// current time will be sent on C, unless the Timer was created by AfterFunc.
func TestTimerBasic(t *testing.T) {
	// A Timer must be created with NewTimer or AfterFunc.
	// A timer stop() function is used to cancel the timer before firing / draining, Stop() returns false if the timer is already drained
	// A timer Reset(Duration d) function is used to reset the time so that it could be used again, calling Reset() on active timer will return true and on drained timer will return false
	
	_10SEC := time.Second * 5
	var timer *time.Timer = time.NewTimer(_10SEC) //a pointer type timer is given
	start := time.Now()
	end := <-timer.C
	t.Logf("waiting for %v sec for this function to execute , start: %v end: %v \n", end.Sub(start).Seconds(), start, end)
	t.Logf("timer.Stop() should return false, got %v \n", timer.Stop())

	//what hapens if the timer.Stop() is called again
	t.Logf("if timer.Stop() called again it should return false, got: %v \n", timer.Stop())
	
	//the timer is drained and not reet what happens if the try to drain from channel
	// t1 := <-timer.C // timer.Stop() will not close the channel. so if we try to drain again without a reset it will be blocking forever.
	// t.Logf("what happens if we try to drain from already drained and stopped channel, drain value : %v \n", t1)

	//timer is reset and we try to drain
	if reset := timer.Reset(_10SEC); reset {
		//reset will be false as it is not called on an active/undrained timer
		t.Logf("test case failed! expecting reset to be false got true")
		t.Fail()
	}
	start = time.Now()
	end = <-timer.C //will wait or 5 sec
	t.Logf("waiting for %v sec for this function to execute , start: %v end: %v \n", end.Sub(start).Seconds(), start, end)

	timer = time.AfterFunc(_10SEC, func(){
		t.Log("this function executed after 5 sec")
	})
	if stopped := timer.Stop(); !stopped {
		t.Log("the 5 sec timer is stopped before firing or execution. so timer should be true, but got false")
		t.Fail()
	}
	//<-timer.C //should wait forever, coz the channel is not closed
}

// This function demonstrates the usage of time.Ticker
// A Ticker holds a channel that delivers “ticks” of a clock at intervals.
// ticker.Stop() Stop turns off a ticker. After Stop, no more ticks will be sent. Stop does not close the channel
// ticker.Reset()  Reset stops a ticker and resets its period to the specified duration. The next tick will arrive after the new period elapses.
func TestTickerBasic(t *testing.T) {
	_2sec, _1sec := time.Second * 2, time.Second
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
	var ticker *time.Ticker = time.NewTicker(_2sec) //NewTIcker gives a pointer
	tickerFunc := func(ctx context.Context) int {
		for {
			select {
			case t1 := <-ticker.C:
				t.Logf("Tick executed at %v", t1.Format(time.Stamp))
			case <-ctx.Done():
				t.Log("quit called. exiting the go routine")
			    return 0
			}	
		}		
	}
	tickerFunc(timeoutContext)
	ticker.Stop()

	//what happens if reset is called after stop, the whole etup will run again
	ticker.Reset(_1sec)
	timeoutContext, cancel = context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
	tickerFunc(timeoutContext)
}	