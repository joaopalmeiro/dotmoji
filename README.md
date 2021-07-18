# dotmoji

A [Go](https://golang.org/) CLI to generate sequences of colored circles with emojis.

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
