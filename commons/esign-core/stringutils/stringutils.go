// Package stringutils provides string manipulation utilities for Squad B services.
// This is a squad-commons module owned by Squad B.
package stringutils

import (
	"strings"
	"unicode"
)

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
	cleaned := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return cleaned == Reverse(cleaned)
}

// WordCount counts the number of words in a string
func WordCount(s string) int {
	fields := strings.Fields(s)
	return len(fields)
}

// Capitalize capitalizes the first letter of each word
func Capitalize(s string) string {
	return strings.Title(strings.ToLower(s))
}

// RemoveSpaces removes all spaces from a string
func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// CountVowels counts the number of vowels in a string
func CountVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

// IsAlphaNumeric checks if a string contains only alphanumeric characters
func IsAlphaNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return len(s) > 0
}

// ReverseWords reverses the order of words in a string
func ReverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
