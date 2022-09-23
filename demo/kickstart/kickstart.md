# Go Bootcamp Kickstart
Kickstart introduces the resources and Concepts and strategy to learn Go. Learning is an experience. We have to try out a lot of probelms to understand the concepts and famirialize with the go syntax an style.

## How to learn Efficiently
- Youtube Videos:  Freecode camp provides a excellent Go introductory course
- WriteDown: Write down what you have learned, Create Cheat sheet
- Code reading: Read code written by experts and start learning new concepts and patterns
- Problem Solving and Coding: Solve problems in LeetCode and CodeChef
- Community Involvement: Involve in Go Slack, Follow Go Forums, Blogs
- Freelance: Freelance as Go programmer in upwork and freelancer
- Commitment: Atleast 2 hrs per day

### Setup
An online [go playground](https://go.dev/play/) is available. It could be used to learn the go basics without any local setup. There are some limitation in the online playground like accessing the file system.
Alternatively, Go can be installed in local machine [install go](https://go.dev/doc/install).
```sh
# check if go is installed
go version

# access go tools help
go help
```

#### Go ENVIRONMENT Variables
The go command and the tools it invokes consult environment variables for configuration. If an environment variable is unset, the go command uses a sensible default setting
```sh
# prints all go environment variables
go env 
```
Refer [Go Environvment variables](https://pkg.go.dev/cmd/go#hdr-Environment_variables) for detailed documentation

**Basic Go Environment varibales**
- GOROOT: The path of the go sdk and libraries
- GOPATH: is used to resolve import statements and dependencies
- GOENV: The location of the Go environment configuration file

**Setup go dev setup**
Visual studio provides excellent support for Go development. Alternatively, any test editir can be used for learning Go

```sh
mkdir -p hello
cd hello
go mod init github.com/sankr888/go-bootcamp #THis will initiale a go module and creates go.mod file

# After module is initialed creat files in hello directory and start coding
```

### HelloWorld, Entrypoint, packages and modules
```go
package main

// fmt is a standard library
// import "fmt" - will also work for single imports
import (
	"fmt"
)

// main entrypoint
func main() {
	fmt.Println("Hello, World!")
}
```
Go Entrypoint is  function named main in package main.
Every source file must belong to a go package. package declaration should be the first line of the source file.Technically Multiple source files in different folders canhave the sam package, But usually it is good to keep sourcefile of same package under same folder. The functions, types and variables should be unique within a package. Anything (variable, type, or function) that starts with a capital letter is exported, and visible outside the package. when u import a package, you can only access its exported names

Modules are collections of unique packages which are built and exported together. All packages in modules have the same prefix name. ex: github.com/sankar888/go-bootcamp The modules should be accessible in the reverse module name url.

```sh
go run fully_qualified_name_of_gofile.go

# will build the create an executable if it has a main entrypoint
go build .

# will build and install the executable in the GOPATH directory along with other 3rd party packages
go install .
```

### Go Learning Resources
1. [Go Official website](https://go.dev/)
2. [Go online playground](https://go.dev/play/)
3. [Go packages](https://pkg.go.dev/)
4. [Go Documentation](https://go.dev/doc/)
5. [Go basics](https://github.com/gophertuts/go-basics/tree/master/packages)
6. [Effective Go Guide](https://go.dev/doc/effective_go)
7. [Go Introductory Freecourseware video](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=348s)

