package main

/*
This line declares that the file is part of the main package.
The main package is a special package in Go that is used to create executable programs.
The Go runtime looks for the main package to start the execution of the program.
*/
import "fmt"

//This line imports the fmt package, which is part of Go's standard library.
//The fmt package provides formatted I/O functions, such as Println used for printing to the console.

func main() {
	fmt.Println("Hello This is a Main package")
	hello()
	new()
	world()

	/* The main function is the entry point of the Go program. When the program starts, the code inside the main function is executed first.
	   fmt.Println("Hello This is a Main package"):
	   This line prints the string "Hello This is a Main package" to the console.
	   hey():
	   This line calls the hey function, which is defined below.
	   bye():
	   This line calls the bye function, which is also defined below.
	   new():
	   This line calls the new function, which is also defined below.
	*/

	/*Its won't works here---->
	    func hello()  {
		fmt.Println("This is a hello package")
	}*/
}
func hello() {
	fmt.Println("This is a hello package")
}
func world() {
	fmt.Println("This is a world package")
}

func new() {
	fmt.Println("This is a new package")
}