package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
* A simple web server which acts as speed tesst callback
 */

func main() {
	printUsage() //todo: add options for usage and help
	port := flag.Int("port", 8080, "The port to server listens to")
	address := flag.String("address", "localhost", "The interface the server listents to")
	flag.Parse()

	fmt.Printf("Starting mockserver on http://%s:%d\n", *address, *port)
	startServer(*address, *port)
}

var usage string = `
mockapi is a cli tool which spins up a mock http server. which accepts all incoming requests and logs it to the console
Usage:
------
mockapi [options]
where options are
--port int - the port to listen to. defaults to 8080
--address string - the ip address to bind to. defaults to localhost

Example:
--------
mockapi --port 8080 --address 10.3.5.12
mockapi //listens on localhost on port 8080
mockapi --port 8080 --address 0.0.0.0 //listens on all interfaces
`

func printUsage() {
	fmt.Println(usage)
}

func startServer(address string, port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("\ninfo:")
		fmt.Printf("%20v:\t%v\n", "time", time.Now())
		fmt.Printf("%20v:\t%v\n", "method", r.Method)
		fmt.Printf("%20v:\t%v\n", "url", r.URL)

		fmt.Println("\nheaders:")
		for key, value := range r.Header {
			fmt.Printf("%20s:\t%v\n", key, value)
		}

		if method := r.Method; method == "POST" || method == "PUT" || method == "PATCH" {
			//print the post body
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println("Error during reading body", err)
			}
			fmt.Println("\nbody:")
			fmt.Println(string(body))
		}

		fmt.Println("\n----------------------------------------------------------------------------")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`success`))
	})
	http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), nil)
}

type httpMethod string

const (
	GET    httpMethod = http.MethodGet
	POST              = http.MethodPost
	PUT               = http.MethodPut
	DELETE            = http.MethodDelete
)

type authType int

const (
	NONE authType = iota
	BASIC
)

type auth struct {
	enabled bool
	atype   authType
	aConfig any
}

type BasicAuthConfig struct {
	username string
	password string
}

func NewBasicAuthConfig(username string, password string) {

}

type mock struct {
	urlPattern string
	method     httpMethod
}

