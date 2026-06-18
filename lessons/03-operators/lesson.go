package lesson03

import (
	"errors"
	"fmt"
)

/*
Lesson 03: Operators
Learning objectives:
- Practice arithmetic, comparison, logical, and bitwise operators.
- Combine operators inside small reusable functions.
- Return errors for invalid operations.
- Review string handling with indexed comparisons.
Estimated time: 20 minutes
Prerequisites: Lessons 01-02.
*/

// Calculate performs a simple arithmetic operation.
func Calculate(a float64, op string, b float64) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}

		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", op)
	}
}

// IsEven uses a bitwise AND to inspect the lowest bit of n.
func IsEven(n int) bool {
	return n&1 == 0
}

// IsPalindrome reports whether s reads the same forwards and backwards.
func IsPalindrome(s string) bool {
	runes := []rune(s)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		if runes[left] != runes[right] {
			return false
		}
	}

	return true
}

// RunDemo prints sample calculations and checks.
func RunDemo() {
	result, err := Calculate(10, "/", 2)
	fmt.Println("10 / 2 =", result, "error:", err)
	fmt.Println("Is 8 even?", IsEven(8))
	fmt.Println("Is level a palindrome?", IsPalindrome("level"))
}

// ---- FOLLOW ALONG ----
// TODO: Add a new case to Calculate for the % operator using integers in a separate function.
// TODO: Test IsEven with an odd number.
// TODO: Try IsPalindrome with a non-palindrome like "gopher".
