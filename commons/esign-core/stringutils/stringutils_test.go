package stringutils

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"Go", "oG"},
		{"", ""},
		{"a", "a"},
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"A man a plan a canal Panama", true},
		{"hello", false},
		{"", true},
		{"a", true},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello world", 2},
		{"", 0},
		{"one", 1},
		{"  multiple   spaces  ", 2},
	}

	for _, test := range tests {
		result := WordCount(test.input)
		if result != test.expected {
			t.Errorf("WordCount(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "Hello World"},
		{"HELLO WORLD", "Hello World"},
		{"", ""},
		{"go", "Go"},
	}

	for _, test := range tests {
		result := Capitalize(test.input)
		if result != test.expected {
			t.Errorf("Capitalize(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 2},
		{"HELLO", 2},
		{"xyz", 0},
		{"", 0},
		{"aeiou", 5},
	}

	for _, test := range tests {
		result := CountVowels(test.input)
		if result != test.expected {
			t.Errorf("CountVowels(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"hello123", true},
		{"hello world", false},
		{"", false},
		{"123", true},
		{"hello!", false},
	}

	for _, test := range tests {
		result := IsAlphaNumeric(test.input)
		if result != test.expected {
			t.Errorf("IsAlphaNumeric(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "world hello"},
		{"Go is awesome", "awesome is Go"},
		{"single", "single"},
		{"", ""},
		{"  leading and trailing  ", "trailing and leading"},
		{"multiple   spaces here", "here spaces multiple"},
	}

	for _, test := range tests {
		result := ReverseWords(test.input)
		if result != test.expected {
			t.Errorf("ReverseWords(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
