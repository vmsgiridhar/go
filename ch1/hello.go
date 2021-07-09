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

	// composite types - array
	var myarr = [3]int{10, 20, 30}
	fmt.Println(myarr)

	// sparse array
	var mysparse = [12]int{10, 20, 6: 30, 40, 50, 60, 70}
	fmt.Println(mysparse)

	// comparing two arrays
	var myarr1 = [...]int{1, 2, 30}
	var myarr2 = [3]int{1, 2, 3}
	fmt.Println("Is myarr1 equal to myarr2:", myarr1 == myarr2)

	// Slices are more important than arrays in Go
	// Using […] makes an array. Using [] makes a slice.
	var myslice = []int{5, 6, 7}
	fmt.Println("The value of myslice is: ", myslice)

	// Multidimensional slice
	var multidslice [2][3]int
	fmt.Println("The Multi dimensional slic is : ", multidslice)
	// The only thing you can compare a slice with is nil:
	// A slice is the first type we’ve seen that isn’t comparable. It is a compile-time error to use == to see if two slices are identical or != to see if they are different.

	// length function on a slice
	var myslice2 []int
	myslice2 = append(myslice2, 10)
	fmt.Println("The length of myslice2 is: ", len(myslice2))

	// slicing the slic
	var myslice3 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var myslice3_x = myslice3[:2]
	fmt.Println("The value of myslice3_x is", myslice3_x)

	// string
	var mystring string = "Hi World!"
	var b byte = mystring[1]
	var bstring string = mystring[0:1]
	fmt.Println("The value of b is: ", b)
	fmt.Println("The value of bstring is: ", bstring)

	// rune, byte and type conversions
	var myrune rune = 's'
	fmt.Println(myrune)

	var mystring_myrune = string(myrune)
	fmt.Println(mystring_myrune)

	var mybyte byte = 'i'
	var mystring_mybyte = string(mybyte)
	fmt.Println(mystring_mybyte)

	// map
	var nilMap map[string]int
	fmt.Println(nilMap)

	mymap := map[string]int{}
	fmt.Println(mymap)

	// map with string as key and an array of strings
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)

	// Using a map
	totalWins := map[string]int{}
	totalWins["India"] = 5
	totalWins["Australia"] = 10
	fmt.Println("Total wins of India are: ", totalWins["India"])

	// delete from a map
	delete(totalWins, "India")
	fmt.Println("TotalWins is now: ", totalWins)

	// As all keys in a map should be the same, we can't use it for an API
	// So, we write a struct
	type person struct {
		name string
		age  int
		pet  string
	}

	// assigning a struct literal
	bob := person{
		"bob",
		24,
		"Tom",
	}

	fmt.Println("The bob is a Person: ", bob)
	fmt.Println("Name of bob is : ", bob.name)
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
// don’t use arrays unless you know the exact length you need ahead of time
