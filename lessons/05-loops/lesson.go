package lesson05

import "fmt"

/*
Lesson 05: Loops
Learning objectives:
- Use for as Go's only loop keyword.
- Practice counted loops and while-style loops.
- Build slices in loops with append.
- Handle input validation before iteration.
Estimated time: 20 minutes
Prerequisites: Lessons 01-04.
*/

// Factorial multiplies every whole number from 1 through n.
func Factorial(n int) (int64, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative numbers do not have factorials")
	}

	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}

	return result, nil
}

// FizzBuzz returns the classic FizzBuzz sequence from 1 to max.
func FizzBuzz(max int) []string {
	if max < 1 {
		return []string{}
	}

	results := make([]string, 0, max)
	for i := 1; i <= max; i++ {
		switch {
		case i%15 == 0:
			results = append(results, "FizzBuzz")
		case i%3 == 0:
			results = append(results, "Fizz")
		case i%5 == 0:
			results = append(results, "Buzz")
		default:
			results = append(results, fmt.Sprintf("%d", i))
		}
	}

	return results
}

// SumDigits adds together every digit in the absolute value of n.
func SumDigits(n int) int {
	if n < 0 {
		n = -n
	}

	sum := 0
	for {
		sum += n % 10
		n /= 10
		if n == 0 {
			break
		}
	}

	return sum
}

// RunDemo shows loop-based examples.
func RunDemo() {
	fact, _ := Factorial(5)
	fmt.Println("5! =", fact)
	fmt.Println("FizzBuzz to 5:", FizzBuzz(5))
	fmt.Println("SumDigits(-123):", SumDigits(-123))
}

// ---- FOLLOW ALONG ----
// TODO: Try Factorial with 0 and 1.
// TODO: Print the FizzBuzz result for 1 through 15.
// TODO: Change SumDigits so it also counts how many digits were processed.
