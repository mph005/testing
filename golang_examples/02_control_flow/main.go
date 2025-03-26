package main

import (
	"fmt"
)

func main() {
	// Conditional statements
	fmt.Println("--- If-Else statements ---")

	// Basic if statement
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	// If-else statement
	y := 3
	if y > 5 {
		fmt.Println("y is greater than 5")
	} else {
		fmt.Println("y is not greater than 5")
	}

	// If with initialization statement
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Switch statements
	fmt.Println("\n--- Switch statements ---")

	// Basic switch
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("End of work week")
	default:
		fmt.Println("Weekend")
	}

	// Switch with no expression (alternative to if-else chains)
	hour := 15
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// Switch with initialization
	switch os := "windows"; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	// Loops
	fmt.Println("\n--- Loops ---")

	// Basic for loop
	fmt.Println("Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For as a while loop
	fmt.Println("For as while loop:")
	counter := 0
	for counter < 5 {
		fmt.Printf("%d ", counter)
		counter++
	}
	fmt.Println()

	// Infinite loop with break
	fmt.Println("Loop with break:")
	sum := 0
	for {
		sum++
		if sum > 5 {
			break
		}
		fmt.Printf("%d ", sum)
	}
	fmt.Println()

	// Loop with continue
	fmt.Println("Loop with continue (even numbers only):")
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue // Skip odd numbers
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Range over a slice
	fmt.Println("Range over slice:")
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i, fruit)
	}

	// Range over a map
	fmt.Println("Range over map:")
	capitals := map[string]string{
		"USA":    "Washington DC",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	for country, capital := range capitals {
		fmt.Printf("%s: %s\n", country, capital)
	}
}
