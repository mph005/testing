// Package advanced provides more complex mathematical functions
package advanced

import "math"

// Square returns the square of a number
func Square(x float64) float64 {
	return x * x
}

// Cube returns the cube of a number
func Cube(x float64) float64 {
	return x * x * x
}

// Sqrt returns the square root of a number
// Returns 0 for negative numbers
func Sqrt(x float64) float64 {
	if x < 0 {
		return 0
	}
	return math.Sqrt(x)
}

// Power returns x raised to the power of y
func Power(x, y float64) float64 {
	return math.Pow(x, y)
} 