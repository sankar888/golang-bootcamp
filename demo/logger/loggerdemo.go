package main

import (
	"log"
	"os"
	"time"

	"github.com/sankar888/golang-bootcamp/demo/common"
)

/*
Go provides a log package, which introduces a type Logger.
Log package also has a standard logger which writes to std.err
log package allows us to set the Writer - destination, configure logging output flages, etc.
It also has helper functions to print formatted output, panic and fatal log messages

Logging Frameworks in go: https://blog.logrocket.com/5-structured-logging-packages-for-go/
*/
func main() {
	defaultLoggerBasicUsage()
	defaultLoggerCustomOptions()
	flagsUsage()
	newLogger()
}

// std logger
func defaultLoggerBasicUsage() {
	common.Start("**Default Go Loger basic usage**")
	defer handlePanic()
	//basic usage
	log.Println("prints this message with default logger and default setings")
	log.Print("log.Print function print the message using default logger")
	log.Print("without a new line ")
	log.Printf("%v, TimeNow: %v", "Prints a formatted log message using default logger", time.Now())

	//panic and fatal
	log.Panic("Panic:What to do ???") //log.Panic raises a panic
	log.Fatalln("Fatal Error")        //This line is not executed as a panic is raised, fatal calls os.exit
	common.End()
}

// std logger with custom options set
func defaultLoggerCustomOptions() {
	common.Start("Default Go Logger - custom options**")
	//log package has helper functions to custamize the behaviour of std logger
	log.Println("default logger msg with default settings")

	//customize default logger
	log.SetPrefix("INFO: ") //prefix to add to all log messages
	log.Println("default logger with prefix set")
	//log.SetOutput() //can set output to a different writer , like a file or socket
	common.End()
}

// logger flag usage
func flagsUsage() {
	/*
	   Log package provides constant integer flags which sets the functionalities of the logger
	   The flags are
	   	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	   	Ltime                         // the time in the local time zone: 01:23:23
	   	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	   	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	   	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	   	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	   	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message

	   Usage:
	   flags Ldate | Ltime (or LstdFlags) produce,
	   2009/01/23 01:23:23 message

	   while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
	   2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	*/
	common.Start("**logger Flag Usage**")
	log.Println("before setting flags")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile | log.LUTC | log.Lmsgprefix) // The order doesn't matter
	log.Println("message from customized logger with all flags set")
	common.End()
}

// we could also create our own logger
func newLogger() {
	common.Start("**New Logger Demo**")
	defer handlePanic()
	flag := log.Ldate | log.Ltime | log.Llongfile | log.LUTC | log.Lmsgprefix
	var logger *log.Logger = log.New(os.Stderr, "bootcamp:", flag) //create a customized new logger

	logger.Println("Hello, From custom Logger")
	logger.Panic("panic: What to do ?")
}

func handlePanic() { //A Panic recovery message which will execute when a panic is raised
	if r := recover(); r != nil {
		log.Printf("Panic: %v,PanicRecovery: Don't worry. I will handle. \n", r)
		common.End()
	}
}
