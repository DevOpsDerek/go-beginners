package lesson09

import (
	"errors"
	"fmt"
)

/*
Lesson 09: Error Handling
Learning objectives:
- Work with the built-in error interface.
- Return custom validation errors.
- Recover from panics in a controlled way.
- Wrap and inspect errors with fmt.Errorf and errors.As.
Estimated time: 25 minutes
Prerequisites: Lessons 01-08.
*/

var errDivisionByZero = errors.New("division by zero")

// ValidationError describes why a field value is invalid.
type ValidationError struct {
	Field   string
	Message string
}

// Error satisfies the error interface.
func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

// SafeDivide divides a by b and returns an error when b is zero.
func SafeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide: %w", errDivisionByZero)
	}

	return a / b, nil
}

// ValidateAge checks whether age falls within a realistic range.
func ValidateAge(age int) error {
	if age < 0 || age > 150 {
		return &ValidationError{Field: "age", Message: "must be between 0 and 150"}
	}

	return nil
}

// SafeRun executes fn and converts a panic into a returned error.
func SafeRun(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	fn()
	return nil
}

// RunDemo shows normal errors and panic recovery.
func RunDemo() {
	result, err := SafeDivide(10, 2)
	fmt.Println("SafeDivide:", result, err)
	fmt.Println("ValidateAge(-1):", ValidateAge(-1))
	fmt.Println("SafeRun panic:", SafeRun(func() { panic("boom") }))
}

// ---- FOLLOW ALONG ----
// TODO: Use errors.As in a test or scratch file to inspect ValidationError.
// TODO: Try SafeDivide with zero.
// TODO: Wrap another sentinel error with fmt.Errorf and %w.
