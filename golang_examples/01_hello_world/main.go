// Package declaration - every Go file must belong to a package
package main

// Import statements for external packages
import (
	"fmt" // The format package for I/O functions
)

// The main function is the entry point of the executable programs
func main() {
	// Print a simple hello world message
	fmt.Println("Hello, World!")

	// Variable declarations
	// Using var keyword with explicit type
	var name string = "Gopher"

	// Short variable declaration with type inference
	age := 5

	// Constants declaration
	const language = "Go"

	// Multiple variable declaration
	var (
		isAwesome bool   = true
		version   string = "1.20"
	)

	// Printing variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Language:", language)
	fmt.Println("Is Awesome:", isAwesome)
	fmt.Println("Version:", version)

	// Basic data types
	var (
		integer int     = 42
		float   float64 = 3.14
		boolean bool    = true
		str     string  = "Go is fun!"
	)

	// Printf for formatted output
	fmt.Printf("Integer: %d\n", integer)
	fmt.Printf("Float: %.2f\n", float)
	fmt.Printf("Boolean: %t\n", boolean)
	fmt.Printf("String: %s\n", str)

	// Zero values (default values when not initialized)
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string

	fmt.Println("\nZero values:")
	fmt.Printf("Default int: %d\n", defaultInt)       // 0
	fmt.Printf("Default float: %.1f\n", defaultFloat) // 0.0
	fmt.Printf("Default bool: %t\n", defaultBool)     // false
	fmt.Printf("Default string: %q\n", defaultString) // ""
}
