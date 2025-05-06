# Go Variables Explained

Variables are like containers that store data in your program.  
In Go, you can declare variables in two main ways:  
1. Using the `var` keyword (explicit declaration)  
2. Using the `:=` syntax (short declaration)

---

## Example Code

```go
// 1. Package declaration
package main

// 2. Import needed packages
import "fmt"

// 3. Main function - where program execution begins
func main() {
    // =============================================
    // METHOD 1: Using 'var' keyword (explicit declaration)
    // =============================================
    
    // Declare a variable of type 'int' (integer)
    // Syntax: var variableName type = value
    var age int = 25
    
    // Declare a variable with type inference
    // Go automatically detects the type based on the value
    var name = "Naruto"
    
    // Declare a variable without initial value
    // It gets the 'zero value' for its type (0 for numbers, "" for strings, false for bool)
    var score int
    
    // =============================================
    // METHOD 2: Short declaration (using :=)
    // =============================================
    
    // Declare and initialize a variable in one line
    // Syntax: variableName := value
    // Go automatically detects the type
    country := "Tokyo"
    
    // Multiple variable declaration
    height, weight := 175.5, 68.2
    
    // =============================================
    // Using the variables
    // =============================================
    score = 95 // Assigning value after declaration
    
    fmt.Println("Name:", name)           // Prints: Name: Naruto
    fmt.Println("Age:", age)             // Prints: Age: 25
    fmt.Println("Country:", country)     // Prints: Country: Tokyo
    fmt.Println("Height:", height)       // Prints: Height: 175.5
    fmt.Println("Weight:", weight)      // Prints: Weight: 68.2
    fmt.Println("Score:", score)         // Prints: Score: 95
    
    // Print variable types
    fmt.Printf("Type of age: %T\n", age)       // %T shows the type
    fmt.Printf("Type of name: %T\n", name)     // Prints: Type of name: string
}
```

---

## Key Concepts Explained

### 1. `var` Keyword Declaration
- Used for **explicit variable declaration**
- Can specify type or let Go infer it
- Can declare without initialization (gets zero value)
- Good for package-level variables or when you need to declare first and assign later

```go
var count int      // Declaration only (zero value = 0)
var price = 19.99  // Type inferred as float64
var message string // Declaration only (zero value = "")
```

### 2. Short Declaration (`:=`)
- Used for **implicit declaration and initialization**
- Type is always inferred
- Must initialize with a value
- Can declare multiple variables at once
- Only works inside functions

```go
username := "sasuke123"  // string
isActive := true         // bool
x, y := 10, 20           // multiple int declarations
```

### 3. Zero Values
When you declare a variable without initialization:
- `0` for numeric types
- `""` for strings
- `false` for booleans
- `nil` for pointers, slices, maps, etc.

### 4. Type Inference
Go automatically detects the type when:
- Using `:=` syntax
- Using `var` with initialization (without explicit type)

### 5. Printing Variables
- `fmt.Println()` - prints with newline
- `fmt.Printf()` - formatted printing (`%T` shows type)

---

## Best Practices
1. Use `:=` inside functions for cleaner code
2. Use `var` at package level or when you need zero values
3. Choose meaningful variable names
4. Initialize variables when possible (avoids zero-value surprises)