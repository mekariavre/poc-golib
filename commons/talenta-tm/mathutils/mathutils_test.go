package mathutils

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(4, 3)
	expected := 12
	if result != expected {
		t.Errorf("Multiply(4, 3) = %d; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	quotient, remainder := Divide(10, 3)
	if quotient != 3 || remainder != 1 {
		t.Errorf("Divide(10, 3) = %d, %d; want 3, 1", quotient, remainder)
	}

	// Test division by zero
	quotient, remainder = Divide(10, 0)
	if quotient != 0 || remainder != 10 {
		t.Errorf("Divide(10, 0) = %d, %d; want 0, 10", quotient, remainder)
	}
}

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	expected := 15
	if result != expected {
		t.Errorf("Sum(%v) = %d; want %d", numbers, result, expected)
	}
}

func TestAverage(t *testing.T) {
	numbers := []int{2, 4, 6}
	result := Average(numbers)
	expected := 4.0
	if result != expected {
		t.Errorf("Average(%v) = %f; want %f", numbers, result, expected)
	}

	// Test empty slice
	empty := []int{}
	result = Average(empty)
	if result != 0 {
		t.Errorf("Average(%v) = %f; want 0", empty, result)
	}
}
func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"positive numbers", []int{1, 2, 3, 4, 5}, 5},
		{"negative numbers", []int{-10, -3, -7, -1}, -1},
		{"mixed numbers", []int{-2, 0, 3, -1, 2}, 3},
		{"single element", []int{42}, 42},
		{"all same", []int{7, 7, 7, 7}, 7},
		{"empty slice", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Max(tt.numbers)
			if result != tt.expected {
				t.Errorf("Max(%v) = %d; want %d", tt.numbers, result, tt.expected)
			}
		})
	}
}
