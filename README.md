# Go Programming: A Beginner-Friendly Roadmap

Welcome to your journey into the Go programming language! This guide is structured in progressive chapters, with a strong focus on practical understanding and test-driven development.

---

## Part 1: Getting Started

### Chapter 1: Welcome to Go & Your First Program

- **What is Go?**
  - A fast, compiled, statically typed language designed by Google.
  - Efficient concurrency, simple syntax, excellent tooling.

- **Why Learn Go?**
  - Popular for web services, command-line tools, and cloud infrastructure.

- **Setting Up Your Go Environment**
  - Installation guides for Windows, macOS, and Linux.
  - Check installation: `go version`

- **Your First Go Program: `Hello, World!`**
  ```go
  package main

  import "fmt"

  func main() {
      fmt.Println("Hello, World!")
  }
  ```
  - `package main` - Entry point for an executable.
  - `import "fmt"` - Imports the standard I/O package.
  - `func main()` - The starting function.
  - `fmt.Println()` - Prints output to the console.

- **Run Your Code**  
  ```bash
  go run hello.go
  ```

- **Intro to Testing**
  - What is testing? It's like checking homework before submission.
  - We’ll test almost everything we build—except `main()` (at least, for now).

---

### Chapter 2: Variables, Basic Types, and Simple Operations

- **Variables**
  - `var` keyword or short `:=` declaration.

- **Basic Types**
  - Integers (`int`)
  - Floating-points (`float64`)
  - Booleans (`bool`)
  - Strings (`string`)
  - Zero values for each type.

- **Simple Operations**
  - Arithmetic: `+`, `-`, `*`, `/`, `%`
  - String concatenation: `+`
  - Comparisons: `==`, `!=`, `<`, `>`, `<=`, `>=`
  - Type conversions: e.g., `float64(x)`
  - Constants: `const`

- **Testing Integration**
  - We'll start writing tests in the next chapter using functions.

---

## Part 2: Building Blocks: Functions and Control Flow

### Chapter 3: Functions: Reusable Code Blocks

- **What Are Functions?**
  - Named blocks of code to reuse logic.

- **Function Syntax**
  ```go
  func Add(a int, b int) int {
      return a + b
  }
  ```

- Parameters, return values, multiple returns, and named returns.

- **Testing Focus**
  - Intro to `testing` package.
  - Writing test files: `add_test.go`
  - Writing test functions:
    ```go
    func TestAdd(t *testing.T) {
        got := Add(2, 3)
        want := 5
        if got != want {
            t.Errorf("got %d, want %d", got, want)
        }
    }
    ```
  - Run tests: `go test`

---

### Chapter 4: Control Flow: Making Decisions and Repeating Actions

- **Conditionals**
  - `if`, `else if`, `else`
  - `switch` statements

- **Loops**
  - `for` as a traditional, `while`, or infinite loop
  - `break`, `continue`

- **Examples**
  - `IsEven(n int) bool`
  - `SumUpTo(n int) int`

- **Testing Focus**
  - Test branches and loops.
  - Example: `TestIsEven`, `TestSumUpTo`

---

## Part 3: Working with Collections of Data

### Chapter 5: Arrays & Slices

- **Arrays**
  - Fixed-size collections.
  - Accessing elements.

- **Slices**
  - Dynamic collections.
  - `make`, literals, `append`, slicing syntax `[low:high]`
  - `len()`, `cap()`
  - Looping with `for` and `range`

- **Examples**
  - `SumSlice(numbers []int) int`
  - `FindFirstNegative(numbers []int) (int, bool)`

- **Testing Focus**
  - Edge cases: empty slices, one-element slices.
  - Example: `TestSumSlice`, `TestFindFirstNegative`

---

### Chapter 6: Maps: Key-Value Collections

- **What Are Maps?**
  - Unordered collections (like dictionaries).

- **Map Operations**
  - `make`, literals
  - Add/update values
  - Key existence: `val, ok := m[key]`
  - Delete keys: `delete(m, key)`
  - Iterate with `range`

- **Examples**
  - `CountCharacters(text string) map[rune]int`
  - `GetUserGreeting(userID int) (string, bool)`

- **Testing Focus**
  - Check key lookup, presence, and logic.
  - Example: `TestCountCharacters`, `TestGetUserGreeting`

---

## Part 4: Structuring Data and Code

### Chapter 7: Structs: Creating Your Own Types

- **What Are Structs?**
  - Groups of fields under one type.

- **Defining Structs**
  ```go
  type Person struct {
      Name string
      Age  int
  }
  ```

- **Methods**
  - Value vs. pointer receivers.
  - `func (p Person) Greet() string { ... }`

- **Example**
  - `Rectangle` with methods `Area()` and `Perimeter()`

- **Testing Focus**
  - Test calculations.
  - Example: `TestRectangleArea`, `TestRectanglePerimeter`

---

### Chapter 8: Pointers: Working with Memory Addresses

- **What Are Pointers?**
  - Variables storing memory addresses.

- **Operators**
  - `&` (address-of), `*` (dereference)

- **Usage in Structs**
  - Efficient struct modifications.

- **Example**
  - `IncreaseAge(p *Person)`

- **Testing Focus**
  - Modifying data via pointers.
  - Handling `nil` safely.
  - Example: `TestIncreaseAge`

---

### Chapter 9: Error Handling: The Go Way

- **Philosophy**
  - Errors are values.

- **Error Handling**
  - `error` interface
  - Check with `if err != nil`
  - Create errors: `errors.New`, `fmt.Errorf`

- **Example**
  ```go
  func Divide(a, b float64) (float64, error) {
      if b == 0 {
          return 0, fmt.Errorf("cannot divide by zero")
      }
      return a / b, nil
  }
  ```

- **Testing Focus**
  - Happy vs. error path.
  - Example: `TestDivide`

---

### Chapter 10: Packages: Organizing Your Code

- **What Are Packages?**
  - A directory of Go files.

- **Creating Your Own**
  - `calculator` package for math functions.
  - Exported (Capitalized) vs. unexported (lowercase) identifiers.
  - Import and use in `main`.

- **Standard Library Overview**
  - `fmt`, `strings`, `math`, `os`, `time`, `errors`, `testing`

- **Testing Focus**
  - Package-specific testing: `go test ./calculator`
  - Example: `calculator_test.go`

---

## Part 5: Improving Your Tests

### Chapter 11: Better Testing Techniques

- **Recap**
  - Basics of test functions and structure.

- **Advanced Testing Techniques**
  - Table-driven tests:
    ```go
    tests := []struct {
        a, b, want int
    }{
        {2, 3, 5},
        {10, -2, 8},
    }
    ```
  - Subtests with `t.Run()`
  - Test helpers for reducing boilerplate
  - Logging: `t.Log()`
  - Code coverage: `go test -cover`

- **Example**
  - Refactor `TestAdd` to table-driven format.

- **Testing Focus**
  - Clean, scalable, and maintainable test design.