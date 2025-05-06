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
	var name = "Alice"

	// Declare a variable without initial value
	// It gets the 'zero value' for its type (0 for numbers, "" for strings, false for bool)
	var score int

	// =============================================
	// METHOD 2: Short declaration (using :=)
	// =============================================

	// Declare and initialize a variable in one line
	// Syntax: variableName := value
	// Go automatically detects the type
	country := "Canada"

	// Multiple variable declaration
	height, weight := 175.5, 68.2

	// =============================================
	// Using the variables
	// =============================================
	score = 95 // Assigning value after declaration

	fmt.Println("Name:", name)       // Prints: Name: Alice
	fmt.Println("Age:", age)         // Prints: Age: 25
	fmt.Println("Country:", country) // Prints: Country: Canada
	fmt.Println("Height:", height)   // Prints: Height: 175.5
	fmt.Println("Weight:", weight)   // Prints: Weight: 68.2
	fmt.Println("Score:", score)     // Prints: Score: 95

	// Print variable types
	fmt.Printf("Type of age: %T\n", age)   // %T shows the type
	fmt.Printf("Type of name: %T\n", name) // Prints: Type of name: string
}
