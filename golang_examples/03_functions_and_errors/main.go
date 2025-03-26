package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Basic function with parameters and return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Function with named return values
func calculate(x, y int) (sum, diff, product int) {
	sum = x + y
	diff = x - y
	product = x * y
	// Naked return - returns the named variables
	return
}

// Variadic function - takes variable number of arguments
func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function that returns a function (closure)
func getCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Function that demonstrates defer
func deferExample() {
	fmt.Println("Start of function")

	// Defer statements are executed in LIFO order (last-in, first-out)
	defer fmt.Println("First defer - executed third")
	defer fmt.Println("Second defer - executed second")
	defer fmt.Println("Third defer - executed first")

	fmt.Println("End of function")
}

// Function demonstrating panic and recover
func panicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("About to panic...")
	panic("something bad happened")
	// Code after panic is not executed
	fmt.Println("This line will never be reached")
}

func main() {
	// Basic function call
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	// Function with error handling
	fmt.Println("\n--- Error handling ---")
	result2, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result2)
	}

	result2, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 0 =", result2) // This won't execute
	}

	// Function with named return values
	s, d, p := calculate(10, 5)
	fmt.Println("\n--- Named returns ---")
	fmt.Println("Sum:", s)
	fmt.Println("Difference:", d)
	fmt.Println("Product:", p)

	// Variadic function
	fmt.Println("\n--- Variadic function ---")
	fmt.Println("Sum of 1, 2, 3:", sumNumbers(1, 2, 3))
	fmt.Println("Sum of 5, 10, 15, 20:", sumNumbers(5, 10, 15, 20))

	// Using a slice with variadic function
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of slice:", sumNumbers(numbers...)) // Note the ... to unpack the slice

	// Closure example
	fmt.Println("\n--- Closure ---")
	counter := getCounter()
	fmt.Println("Counter:", counter()) // 1
	fmt.Println("Counter:", counter()) // 2
	fmt.Println("Counter:", counter()) // 3

	// Create a new counter
	counter2 := getCounter()
	fmt.Println("Counter2:", counter2()) // 1 (separate instance)

	// Defer example
	fmt.Println("\n--- Defer ---")
	deferExample()

	// Panic and recover example
	fmt.Println("\n--- Panic and Recover ---")
	panicAndRecover()

	// Common error handling pattern in Go
	fmt.Println("\n--- Typical Error Handling ---")
	_, err = os.Open("non-existent-file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	// Type assertion for specific error types
	_, err = strconv.Atoi("not-a-number")
	if err != nil {
		fmt.Println("Conversion error:", err)
		// Check specific error types
		if _, ok := err.(*strconv.NumError); ok {
			fmt.Println("This is a numeric conversion error")
		}
	}
}
