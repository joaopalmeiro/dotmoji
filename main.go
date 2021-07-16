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

	// Printf (formatted strings).
	// Add `\n` so that `%` does not appear
	// at the end of the string on the command
	// line (https://stackoverflow.com/a/22816721)

	// `%v` (format specifier): Default format.
	fmt.Printf("Message: %v\n", message)

	// `%q`: Put quotes around the variable (for strings).
	fmt.Printf("Message: %q\n", message)

	// `%T`: Get the variable type.
	fmt.Printf("Message: %T\n", message)
}
