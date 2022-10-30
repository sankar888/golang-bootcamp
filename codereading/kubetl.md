# Introduction
Kubectl - client of kubernetes is built using go. undrstand how kubectl client code is written. Understand and discover the various patterns used to build kubernetes client.

## Study
    
### **1** what kubectl comand invokes. Is it a single binary or a shell script ?
Yes. It is a direct binary, possibly built using go. so when we invoke kubectl version or kubectl --help, it might call the main function. How are these arguments (or) options passed to main function.

### **2.** Where is the main function of kubernetes module ?
The main function of kubectl module is present in kubernetes repo. under kubernetes/cmd/kubectl/kubectl.go

### **3.** How does the main function gets the arguments (or) commands?
It seems the main function creates a kubernetes cmd provuided by ```k8s.io/kubectl/pkg/cmd``` and the command is executed by ```k8s.io/component-base/cli```. The cmd 

## ToKnow
1. Learn about the buildt -in go args parsing utility.
2. Learn how to use [cobra](https://github.com/spf13/cobra) go library

## TODO
1. Create a simple command line utility using in-built go utilities.
2. Use cobra library and create a command line utility

# Resources
1. [kubectl repo](https://github.com/kubernetes/kubectl)
2. [Kubernetes org repo](https://github.com/kubernetes)
3. [kubectl main function](https://github.com/kubernetes/kubernetes/blob/master/cmd/kubectl/kubectl.go)
4. [cobra cmd line library](https://github.com/spf13/cobra)