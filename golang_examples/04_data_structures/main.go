package main

import (
	"fmt"
)

func main() {
	// Arrays
	fmt.Println("--- Arrays ---")

	// Fixed-size array with explicit size
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println("Colors array:", colors)

	// Array initialization
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers array:", numbers)

	// Array with size determined by initialization
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println("Days array:", days)
	fmt.Println("Length of days array:", len(days))

	// Slices
	fmt.Println("\n--- Slices ---")

	// Create a slice from an array
	weekdays := days[0:3] // Monday, Tuesday, Wednesday
	fmt.Println("Weekdays slice:", weekdays)

	// Slices are references to underlying arrays
	weekdays[0] = "MONDAY"
	fmt.Println("Modified weekdays slice:", weekdays)
	fmt.Println("Original days array (also modified):", days)

	// Creating slices directly
	fruits := []string{"Apple", "Banana", "Orange"}
	fmt.Println("Fruits slice:", fruits)

	// Slice with make function (allocates a zeroed array and returns a slice)
	// make([]T, length, capacity)
	scores := make([]int, 3, 5)
	fmt.Println("Scores slice:", scores)  // [0 0 0]
	fmt.Println("Length:", len(scores))   // 3
	fmt.Println("Capacity:", cap(scores)) // 5

	// Appending to slices
	scores = append(scores, 100)
	scores = append(scores, 200)
	fmt.Println("After append:", scores) // [0 0 0 100 200]

	// Append beyond capacity - automatically creates a new underlying array
	scores = append(scores, 300)
	fmt.Println("After append beyond capacity:", scores)
	fmt.Println("New capacity:", cap(scores)) // Capacity increases

	// Slice expressions
	// slice[low:high] - items from low to high-1
	// slice[low:] - items from low to end
	// slice[:high] - items from beginning to high-1
	// slice[:] - all items
	allScores := scores[:]
	someScores := scores[2:5]
	fmt.Println("All scores:", allScores)
	fmt.Println("Some scores:", someScores)

	// Maps
	fmt.Println("\n--- Maps ---")

	// Declare a map
	var countryCapitals map[string]string

	// Initialize with make
	countryCapitals = make(map[string]string)

	// Add key-value pairs
	countryCapitals["USA"] = "Washington D.C."
	countryCapitals["India"] = "New Delhi"
	countryCapitals["Japan"] = "Tokyo"

	fmt.Println("Country capitals map:", countryCapitals)

	// Map literal
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 22,
	}
	fmt.Println("Ages map:", ages)

	// Accessing values
	aliceAge := ages["Alice"]
	fmt.Println("Alice's age:", aliceAge)

	// Check if key exists
	bobAge, exists := ages["Bob"]
	if exists {
		fmt.Println("Bob's age:", bobAge)
	}

	// Key doesn't exist
	daveAge, exists := ages["Dave"]
	if !exists {
		fmt.Println("Dave not found in the map, zero value returned:", daveAge)
	}

	// Delete a key
	delete(ages, "Carol")
	fmt.Println("After deleting Carol:", ages)

	// Struct
	fmt.Println("\n--- Structs ---")

	// Define a struct type
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	// Create a struct instance
	var p1 Person
	p1.FirstName = "John"
	p1.LastName = "Doe"
	p1.Age = 30

	fmt.Println("Person 1:", p1)

	// Struct literal
	p2 := Person{
		FirstName: "Jane",
		LastName:  "Smith",
		Age:       25,
	}
	fmt.Println("Person 2:", p2)

	// Struct literal without field names (order matters)
	p3 := Person{"Alice", "Johnson", 22}
	fmt.Println("Person 3:", p3)

	// Nested structs
	type Address struct {
		Street  string
		City    string
		Country string
	}

	type Employee struct {
		Person  // Embedded struct (anonymous field)
		Address Address
		Salary  float64
	}

	emp := Employee{
		Person: Person{
			FirstName: "Bob",
			LastName:  "Brown",
			Age:       40,
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "Boston",
			Country: "USA",
		},
		Salary: 60000.0,
	}

	fmt.Println("Employee:", emp)

	// Access fields of embedded struct
	fmt.Println("Employee name:", emp.FirstName, emp.LastName) // Promoted fields
	fmt.Println("Employee address:", emp.Address.Street, emp.Address.City)
}
