package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

// TODO: add debug or verbose argument.
func main() {
	args := input()
	err := serve(args)
	if err != nil {
		log.Fatal(err)
	}
}

type serverArgs struct {
	address string
	port    int
}

func input() serverArgs {
	address := flag.String("address", "", "the ip address the server binds to. use 0.0.0.0 to bind on all interfaces.")
	port := flag.Int("port", 8080, "the port the server listens to.")
	flag.Parse()

	args := serverArgs{
		address: *address,
		port:    *port,
	}
	if valid, errs := args.validate(); !valid {
		log.Fatalf("invalid arguments! %v", errs)
	}
	return args
}

func (args serverArgs) validate() (valid bool, errs []error) {
	valid, errs = false, make([]error, 0)
	if len(args.address) == 0 || net.ParseIP(args.address) == nil {
		errs = append(errs, errors.New("invalid ip address"))
	}
	if args.port < 1 || args.port > 65535 {
		errs = append(errs, errors.New("the specified port is not within valid port range [1-65535]"))
	}
	if len(errs) == 0 {
		valid = true
	}
	return valid, errs
}

func (args serverArgs) getAddrWithPort() string {
	return fmt.Sprintf("%s:%d", args.address, args.port)
}

func serve(args serverArgs) error {
	server := &http.Server{
		Addr:    args.getAddrWithPort(),
		Handler: bindings(),
	}
	log.Printf("server listening on %s", server.Addr)
	return server.ListenAndServe()
}

func bindings() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RequestURI() != "/" {
			http.NotFound(w, r)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`Welcome to json Beautifier`))
		}
	})

	// serveMux.handleFunc("/json/beautify", func(w http.ResponseWriter, r *http.Request) {

	// })

	// serveMux.handleFunc("/json/minify", func(w http.ResponseWriter, r *http.Request) {

	// })
	return serveMux
}
