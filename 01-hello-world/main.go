package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

/* How to run code inside project: go run main.go
-- GO CLI --
go build: compiles a bunch of go source codes
go run: compiles and executes one or two files
go fmt: formats all the code in each file in the current directory
go install: compiles and 'instals' a package
go get: download the raw source code of someone else's package
go test: runs any tests associated with the current project

What is 'package main' mean: Defines a package that can be compiled and then 'executed'. Must have a func called 'main'
The package main is a private word, and the go build only create an .exe for this word, if the name differs from main, there's no .exe created when go build executes
-- Types of packages --
|-Executable: generates a file that we can run
|-Reusable: code used as 'helpers'. Good place to put reusable logic
|-- Package calculator: Defines a package that can be used as dependency (helper code)
|-- Package uploader: Defines a package that can be used as dependency (helper code)

What does 'import "fmt"' mean: import the fmt package, that implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs'are derived from C's but are simpler

How is the main.go file organized:
Very top: package declaration
Bellow package declaration: import packages
Bellow imports: Funcs, variables etc

*/
