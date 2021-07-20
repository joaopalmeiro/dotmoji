# dotmoji

A [Go](https://golang.org/) CLI to generate emoji sequences.

## References

- The Net Ninja ([repo](https://github.com/iamshaunjp/golang-tutorials)):
  - [Go (Golang) Tutorial #1 - Introduction & Setup](https://youtu.be/etSN4X_fCnM).
  - [Go (Golang) Tutorial #2 - Your First Go File](https://youtu.be/RI9ngRqn9N4).
  - [Go (Golang) Tutorial #3 - Variables, Strings & Numbers](https://youtu.be/98tQPM3R3qU).
  - [Go (Golang) Tutorial #4 - Printing & Formatting Strings](https://youtu.be/m1Uy0WQ2Xns).
  - [Go (Golang) Tutorial #5 - Arrays & Slices](https://youtu.be/Arb-LjPg7FA).
  - [Go (Golang) Tutorial #6 - The Standard Library](https://youtu.be/BoooHY57RXw).
  - [Go (Golang) Tutorial #7 - Loops](https://youtu.be/CL13xV2dHCg).
  - [Go (Golang) Tutorial #8 - Booleans & Conditionals](https://youtu.be/d5oUb2T2iCE).
  - [Go (Golang) Tutorial #9 - Using Functions](https://youtu.be/X68JmClzap4).
  - [Go (Golang) Tutorial #10 - Multiple Return Values](https://youtu.be/ypV7r1ODZCA).
  - [Go (Golang) Tutorial #11 - Package Scope](https://youtu.be/XYK4Rs80q6c).
  - [Go (Golang) Tutorial #12 - Maps](https://youtu.be/v3RodjH1i6c).
  - [Go (Golang) Tutorial #13 - Pass By Value](https://youtu.be/LBLN4Wu5U4w).
  - [Go (Golang) Tutorial #14 - Pointers](https://youtu.be/4B2rwYvuiBo).
  - [Go (Golang) Tutorial #15 - Structs & Custom Types](https://youtu.be/zfMhSFpBSaw).
  - [Go (Golang) Tutorial #16 - Receiver Functions](https://youtu.be/HE6tbWlymmk).
  - [Go (Golang) Tutorial #17 - Receiver Functions with Pointers](https://youtu.be/cgBA5k50At8).
  - [Go (Golang) Tutorial #18 - User Input](https://youtu.be/20HlPtQuc_g).
  - [Go (Golang) Tutorial #19 - Switch Statements](https://youtu.be/l6zlJFJst8o).
  - [Go (Golang) Tutorial #20 - Parsing Floats](https://youtu.be/ipUWKRPw43I).
  - [Go (Golang) Tutorial #21 - Saving Files](https://youtu.be/J88pG3NeRoA).
  - [Go (Golang) Tutorial #22 - Interfaces](https://youtu.be/lbW-KVdIXaY).

## Notes

- [Circles](https://emojipedia.org/search/?q=circle).
- Go is an object-oriented language (in its own way).
- [Uninstall Go](https://golang.org/doc/manage-install#uninstalling).
- [VS Code Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go).
- Commands:
  - `go version` (Go 1.16).
  - `go run main.go`.
  - `go run main.go utils.go`.
- [Numeric types](https://golang.org/ref/spec#Numeric_types) and [types](https://pkg.go.dev/builtin#pkg-types).
- [`fmt`](https://pkg.go.dev/fmt):
  - `%f`: For floats.
  - `%0.1f`: One decimal place.
- [Standard library](https://pkg.go.dev/std) (built-in packages).
- Go is a pass-by-value language (it creates "copies" of values when passed into functions).
- Non-pointer values:
  - Strings.
  - Ints.
  - Floats.
  - Booleans.
  - Arrays.
  - Structs.
- Pointer wrapper values:
  - Slices.
  - Maps.
  - Functions.
- [Pointers](https://tour.golang.org/moretypes/1):
  - A pointer holds the memory address of a value.
  - `&`: pointer.
  - `*`: pointer's underlying value.
  - Use an asterisk (`*`) before the type of an argument in a function for pointers.
- [Switch documentation](https://golang.org/ref/spec#Switch_statements).
- `panic(err)` ([source](https://gobyexample.com/panic)): to abort if a function returns an error that we don't want or know how to handle.
- An interface groups types based on their methods ([example](https://gobyexample.com/interfaces)).
- Toggle `Use Regular Expression` in the search bar in VS Code to search for `\t` ([source](https://github.com/Microsoft/vscode/issues/36321#issuecomment-336790788)).

### Course file (`main.go`)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

// Can be used as a type for function arguments, for example.
type shape interface {
    area() float64
    perimeter() float64
}

// Custom type/blueprint.
type person struct {
    name string
    age  int
}

func newPerson(name string, age int) person {
    p := person{
        name: name,
        age:  age,
    }

    return p
}

// Receiver function.
// `(p person)` or `(p *person)`.
func (p person) format() string {
    return fmt.Sprintf("Person: %v (%v years old)", p.name, p.age)
}

// Receiver function with pointer.
// For structs, pointers are automatically dereferenced.
func (p *person) updateAge(age int) {
    // (*p).age = age
    // or
    p.age = age
}

func sayHello(n string) {
    fmt.Printf("Hello, %v!\n", n)
}

// `string` -> `(string, string)` if multiple return values.
func prepareHello(n string) string {
    return fmt.Sprintf("Hello, %v!", n)
}

func cycleNames(n []string, f func(string)) {
    for _, v := range n {
        f(v)
    }
}

// Helper function.
// `reader := bufio.NewReader(os.Stdin)`.
func getInput(prompt string, r *bufio.Reader) (string, error) {
    fmt.Print(prompt)

    input, err := r.ReadString('\n')

    return strings.TrimSpace(input), err
}

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

    // `for` is also used for while loops.
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // If one of the two is not needed, use `_`.
    for index, value := range numbers {
        fmt.Println(index, value)
    }

    // `if`, `else if`, and `else`.
    // `continue` and `break`.

    sayHello("World")
    fmt.Println(prepareHello("World"))

    cycleNames([]string{"John", "Mary"}, sayHello)

    // Maps: Similar to Python dictionaries.
    // `map[<keys>]<values>`.
    itemPrices := map[string]float64{
        "item1": 4.99,
        "item2": 7.99,
    }

    fmt.Println(itemPrices)
    fmt.Println(itemPrices["item1"])

    for key, value := range itemPrices {
        fmt.Println(key, value)
    }

    myPerson := newPerson("João", 25)

    fmt.Println(myPerson)
    fmt.Println(myPerson.format())

    myPerson.updateAge(26)
    fmt.Println(myPerson.format())

    num, err := strconv.ParseFloat("25.5", 64)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Number (%T): %0.1f\n", num, num)
    }

    // Save file.
    os.WriteFile("message.txt", []byte(message), 0644)
}
```
