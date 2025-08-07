// Package mathutils provides basic mathematical operations for Squad A services.
// This is a squad-commons module owned by Squad A.
package mathutils

// Add performs addition of two integers
func Add(a, b int) int {
	return a + b
}

// Subtract performs subtraction of two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply performs multiplication of two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide performs division of two integers, returns result and remainder
func Divide(a, b int) (quotient, remainder int) {
	if b == 0 {
		return 0, a
	}
	return a / b, a % b
}

// Sum calculates the sum of a slice of integers
func Sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Average calculates the average of a slice of integers
func Average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return float64(Sum(numbers)) / float64(len(numbers))
}

// Max returns the maximum value in a slice of integers
func Max(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}
