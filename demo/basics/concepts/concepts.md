# Common GO Environment variables

    GOROOT - Specifies the root directory of the Go installation.
    GOPATH - Specifies the workspace location. This is where your Go projects and dependencies reside.
    GOBIN - Specifies the directory where binaries are installed when you run go install.
    GO111MODULE - Controls the module mode. Can be on, off, or auto.
    GOPROXY - Specifies the proxy URL for downloading Go modules.
    GOSUMDB - Specifies the checksum database to use for verifying module integrity.
    GOINSECURE - Specifies module paths to be treated as insecure.
    GOPRIVATE - Specifies module paths to be treated as private and not to use the checksum database.
    GOARCH - Specifies the target architecture for the build.
    GOOS - Specifies the target operating system for the build.
    CGO_ENABLED - Controls whether the Cgo tool is enabled. Can be 0 (disabled) or 1 (enabled).

# Modules

Go modules are a dependency management system introduced in Go 1.11 and became the default in Go 1.13. 
They provide a standardized way to manage dependencies, versioning, and package distribution in Go projects.
```shell
# Initialize a Module:
go mod init example.com/mymodule

# Add Dependencies:
go get github.com/some/dependency@v1.2.3

# Build the Module:
go build
```
Go modules is a collection of packages whose lifecycle is managed together. it is similar to a maven project.
Go modules provides a convenient way to manage dependencies. **go.mod** in the root directory mark the collection of packages as a module. 

# Packages
A package is a collection of related Go source files located in the same directory. 
Each source file must start with a package declaration, indicating the package to which it belongs. Package concept helps to
- Organize and manage code.
- Promote code reuse.
- Encapsulate implementation details.
- Manage namespaces and prevent naming conflicts.
- Support modular programming.
- Handle dependencies and versioning.
  
By convention, packages are given lower case, single-word names; there should be no need for underscores or mixedCaps
Also, it is a convention and best practice for a package to have a directory with the same name as the package itself.

