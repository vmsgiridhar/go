package main

import (
	"fmt"
)

func main() { // this opening brace always be at where the function is declared. Writing '{' in the next line will break as the go will treat main() as main();
	fmt.Println("Hello, world!")
	// String with backquote
	fmt.Println(`Hello "World!"`)

	// Data types in Go - Boolean
	var tf bool
	fmt.Println("The value of tf is:", tf)

	// Data types in Go - Numeric (12 types)
	/*
		int8, int16, int32, int64,
		uint8, uint16, uint32, uint64
		byte is uint8
		int is uint32 or unint64 depending on PC
		amd64p32, mips64p32, and mips64p32le
		rune,uintptr
	*/
	var i8 int8 = 127
	fmt.Println("The value in i8 is: ", i8)

	// just a basic math operation
	var x int = 10
	x = x + 10
	fmt.Println("The value of x now is: ", x)

	// floating point types in Go
	/*
		IEEE 754 is followed for storing float values.
			float32 -
			float64 -

		complex64, complex128 are also there.
	*/
	var mypi float64 = 3.14
	fmt.Println("The value of mypi/3 is:", mypi/3)

	// Declaring variables
	var xis0 int

	var xy = 5

	zc := 10

	var (
		xz int    = 1
		y         = 2
		z  int    = 3
		ab string = "ab" + "cd"
	)
	fmt.Println("Different ways of declaring variables")
	fmt.Println(xis0, xy, xz, y, z, ab, zc)

	// const Constants
	const myconst string = "My Constant"
	fmt.Println(myconst)
}

// Use go run when you want to treat a Go program like a script and run the source code immediately.
// go run hello.go will run and give Hello, world! as an output.
// go is a compiled language, but when we execute go run it just creates a temporary binary in a temporary folder that gets deleted soon after program executes.
// go build hello.go builds a binary for later usage instead
// go build hello.go builds an executable in the current directory.
// go build -o hello_world hello.go creates an executable names hello_world in whatever destination mentioned.
// go build -o another_executable/hello_world hello.go will create hello_world executable under another_executable.
// Always run go fmt or goimports before compiling your code!
// Linting and Vetting will solve some of the code styling problems
/*
Make golint and go vet (or golangci-lint) part of your development process to avoid common bugs and nonidiomatic code. But if you are using golangci-lint, make sure your team agrees on the rules that you want to enforce!
*/
// It is a compile-time error to declare a local variable and to not read its value.
