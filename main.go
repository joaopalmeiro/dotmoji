package main

import (
	"fmt"
	"sort"
	"strings"
)

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

	// Sprintf (save formatted strings).
	var longMessage string = fmt.Sprintf("Message: %v", message)
	fmt.Println(longMessage)

	// Array (fixed length).
	// var numbers [3]int = [3]int{1, 2, 3}
	// var numbers = [3]int{1, 2, 3}
	numbers := [3]int{1, 2, 3}
	numbers[1] = 4

	fmt.Println(numbers, len(numbers))

	// Slices (closer to JavaScript arrays, for example).
	var scores = []int{100, 200, 300}
	scores[1] = 50

	// It is not possible to do this with arrays.
	scores = append(scores, 500)

	fmt.Println(scores, len(scores))

	// Slice ranges.
	rangeNumbers := numbers[1:3]
	fmt.Println(rangeNumbers)

	// strings package.
	fmt.Println(strings.Contains(message, "Hello"))
	fmt.Println(strings.ReplaceAll(message, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(message))
	fmt.Println(strings.Index(message, "ll")) // Position.
	fmt.Println(strings.Split(message, " "))

	// sort package.
	// `sort.Strings()` for strings.
	sort.Ints(scores) // In-place.
	fmt.Println(scores)

	// `sort.SearchStrings()` for strings.
	fmt.Println(sort.SearchInts(scores, 50)) // Position.
}
