package main

import "fmt"

// Application entry point.
func main() {
	// Type (`string`) can be omitted.
	var message string = "Hello, World!"

	// var message string

	// Shorthand.
	// Cannot be used outside of a function.
	// message := "Hello, World!"

	// fmt.Print(message)
	fmt.Println(message)
}
