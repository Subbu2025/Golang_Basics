package main

//This line declares that the file is part of the main package.

// file scope-Indicates that this import statement is visible throughout the entire file.
import "fmt"

//The fmt package provides formatted I/O functions, such as Println used for printing to the console.

// package scope-Indicates that this constant is visible throughout the entire main package, which includes all functions and methods within this file.
const ok = true

var amount = 10000

// package scope-Indicates that the main function is visible throughout the main package
func main() { //block scope starts-Indicates that the curly brace { marks the start of the function's block scope.
	fmt.Println("This is for scopes")
	// hello and ok are visible here
	var hello = "Hello"
	fmt.Println(hello, ok)
	fmt.Println(amount, "--Here variable 'amount' has main scope")
	hey()
	fmt.Println(amount, "--Here again variable 'amount' has main scope")
} // block scope ends

func hey() {
	//fmt.Println(hello)---It won't works because variable hello is limitted to main() function block only
	fmt.Println(ok, "--This constant has package scope") //It Works because constant 'ok' is having package scope, so it's visible throughout the main package
	var amount = 5000
	fmt.Println(amount, "--Here variable 'amount' has nested scope")
}