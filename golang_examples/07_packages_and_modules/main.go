package main

import (
	"fmt"

	// Import our local packages using the module path
	"github.com/example/golang_examples/07_packages_and_modules/calculator"
	"github.com/example/golang_examples/07_packages_and_modules/calculator/advanced"
)

func main() {
	fmt.Println("--- Package Example ---")

	// Using the calculator package
	a, b := 10.0, 5.0

	fmt.Printf("Basic operations on %.1f and %.1f:\n", a, b)
	fmt.Printf("Add: %.1f\n", calculator.Add(a, b))
	fmt.Printf("Subtract: %.1f\n", calculator.Subtract(a, b))
	fmt.Printf("Multiply: %.1f\n", calculator.Multiply(a, b))
	fmt.Printf("Divide: %.1f\n", calculator.Divide(a, b))

	// Using the advanced subpackage
	x := 4.0
	fmt.Printf("\nAdvanced operations on %.1f:\n", x)
	fmt.Printf("Square: %.1f\n", advanced.Square(x))
	fmt.Printf("Cube: %.1f\n", advanced.Cube(x))
	fmt.Printf("Square Root: %.1f\n", advanced.Sqrt(x))
	fmt.Printf("Power of 3: %.1f\n", advanced.Power(x, 3))

	// Combining packages
	result := advanced.Square(calculator.Add(a, b))
	fmt.Printf("\nSquare of (%.1f + %.1f) = %.1f\n", a, b, result)
}

/*
Package Organization Notes:

1. Package Declaration:
   - Every Go file must start with a package declaration
   - Files in the same directory share the same package name
   - The package name is typically the same as the directory name

2. Module:
   - A module is a collection of related packages
   - Defined by a go.mod file at the root
   - Contains the module path and Go version
   - Can specify dependencies on other modules

3. Importing Packages:
   - Use the import statement to access other packages
   - Standard library packages can be imported directly
   - Local packages are imported with their module path
   - Third-party packages are imported with their full module path

4. Visibility:
   - Names beginning with uppercase letters are exported (public)
   - Names beginning with lowercase letters are unexported (private)
   - Only exported names can be accessed from outside the package

5. Package Naming Conventions:
   - Use short, concise, lowercase names
   - Avoid underscores or mixedCaps
   - Package name should be the same as the last element of the import path

6. Package Organization:
   - Each package should have a focused purpose
   - Group related functionality into a single package
   - Break large packages into smaller, more focused ones
*/
