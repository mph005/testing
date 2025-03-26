package main

import (
	"fmt"
	"math"
	"strings"
)

// Define a struct type
type Rectangle struct {
	Width  float64
	Height float64
}

// Method with a receiver
// (r Rectangle) is the receiver - this method is associated with Rectangle type
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method with a pointer receiver
// (r *Rectangle) pointer receiver allows the method to modify the Rectangle
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Another struct type
type Circle struct {
	Radius float64
}

// Method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Interface definition
// Any type that implements all the methods of an interface
// is said to implement that interface
type Shape interface {
	Area() float64
}

// Function that accepts an interface
func PrintArea(s Shape) {
	fmt.Printf("Area of shape: %.2f\n", s.Area())
}

// Another interface
type Scaler interface {
	Scale(factor float64)
}

// Multiple interfaces can be combined
type ScalableShape interface {
	Shape
	Scaler
}

// Empty interface can hold values of any type
func PrintInfo(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

// Custom string formatter
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Implement the Stringer interface (from fmt package)
// This is used when the type is formatted with %s or %v
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// Type with multiple interface implementations
type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c Counter) String() string {
	return fmt.Sprintf("Count: %d", c.count)
}

func (c Counter) Count() int {
	return c.count
}

// Interface for counting
type Counting interface {
	Count() int
}

// String slice type for sorting example
type StringSlice []string

// Len returns the length of the slice
func (s StringSlice) Len() int {
	return len(s)
}

// Less compares two elements
func (s StringSlice) Less(i, j int) bool {
	return strings.ToLower(s[i]) < strings.ToLower(s[j])
}

// Swap exchanges two elements
func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	// Methods
	fmt.Println("--- Methods ---")

	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())

	// Pointer receiver method
	rect.Scale(2) // Go automatically handles the conversion to pointer
	fmt.Printf("After scaling: %+v\n", rect)
	fmt.Printf("New area: %.2f\n", rect.Area())

	// Another type with same method
	circle := Circle{Radius: 2}
	fmt.Printf("Circle area: %.2f\n", circle.Area())

	// Interfaces
	fmt.Println("\n--- Interfaces ---")

	// Both Rectangle and Circle implement the Shape interface
	// We can use them wherever a Shape is expected
	PrintArea(rect)
	PrintArea(circle)

	// Slice of Shapes containing different types
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 5},
		Rectangle{Width: 7, Height: 2},
	}

	for _, shape := range shapes {
		PrintArea(shape)
	}

	// Type assertions and type switches
	fmt.Println("\n--- Type Assertions ---")

	for _, shape := range shapes {
		// Type assertion to check if shape is a Rectangle
		if rect, ok := shape.(Rectangle); ok {
			fmt.Printf("Found rectangle: %+v\n", rect)
		}

		// Type switch
		switch s := shape.(type) {
		case Rectangle:
			fmt.Printf("Rectangle: %+v\n", s)
		case Circle:
			fmt.Printf("Circle with radius: %.2f\n", s.Radius)
		default:
			fmt.Printf("Unknown shape: %T\n", s)
		}
	}

	// Empty interface
	fmt.Println("\n--- Empty Interface ---")

	// Empty interface can hold any value
	var anything interface{}

	anything = 42
	PrintInfo(anything)

	anything = "Hello, Go"
	PrintInfo(anything)

	anything = []string{"apple", "banana", "cherry"}
	PrintInfo(anything)

	// String representation
	fmt.Println("\n--- Custom String Formatting ---")

	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Stringer interface is used by fmt package
	fmt.Println(person) // Will use the String() method
	fmt.Printf("%s\n", person)

	// Interface composition
	fmt.Println("\n--- Interface Composition ---")

	counter := &Counter{count: 0}
	counter.Increment()
	counter.Increment()
	fmt.Println(counter) // Uses String() method

	var counting Counting = counter
	fmt.Printf("Counter value: %d\n", counting.Count())

	// Practical interface example: sort package
	fmt.Println("\n--- Practical Example: Sorting ---")

	// Using the sort.Interface from the sort package
	// sort.Interface requires: Len(), Less(i, j), Swap(i, j)
	fruits := StringSlice{"banana", "Apple", "cherry", "Date"}
	fmt.Println("Original slice:", fruits)

	// Instead of calling sort.Sort, we'll show the interface principle
	// by manually sorting using the "bubble sort" algorithm
	for i := 0; i < fruits.Len()-1; i++ {
		for j := 0; j < fruits.Len()-i-1; j++ {
			if fruits.Less(j+1, j) {
				fruits.Swap(j, j+1)
			}
		}
	}

	fmt.Println("Sorted slice:", fruits)
}
