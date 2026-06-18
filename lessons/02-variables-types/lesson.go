package lesson02

import "fmt"

/*
Lesson 02: Variables & Types
Learning objectives:
- Use var declarations, short declarations, and constants.
- Understand zero values and basic built-in types.
- Practice simple type conversion and formatting.
- Use iota-style thinking for related constants.
Estimated time: 20 minutes
Prerequisites: Lesson 01 or basic familiarity with Go functions.
*/

const (
	// These constants show how iota can create a sequence of related values.
	levelBeginner = iota
	levelPracticing
	levelConfident
)

// CelsiusToFahrenheit converts a Celsius temperature to Fahrenheit.
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// TypeName reports the dynamic type of a value using fmt.Sprintf and %T.
func TypeName(v any) string {
	return fmt.Sprintf("%T", v)
}

// AbsoluteValue returns the non-negative version of n without using math.Abs.
func AbsoluteValue(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

// RunDemo shows examples of the lesson functions.
func RunDemo() {
	var city string
	age := 3
	const language = "Go"

	fmt.Println("Zero value example:", city)
	city = "Seattle"
	fmt.Println("Assigned value example:", city)
	fmt.Println("Short declaration example:", age)
	fmt.Println("Constant example:", language)
	fmt.Println("Temperature:", CelsiusToFahrenheit(20))
	fmt.Println("Type name:", TypeName(true))
	fmt.Println("Absolute value:", AbsoluteValue(-42))
	fmt.Println("iota example level:", levelPracticing)
}

// ---- FOLLOW ALONG ----
// TODO: Declare your own constant for a favorite number.
// TODO: Convert 100 Celsius to Fahrenheit.
// TODO: Call TypeName with a string, int, and bool.
