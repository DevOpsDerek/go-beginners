package lesson05

import (
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	t.Run("factorial of five", func(t *testing.T) {
		got, err := Factorial(5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 120 {
			t.Errorf("got %d, want 120", got)
		}
	})

	t.Run("factorial of zero", func(t *testing.T) {
		got, err := Factorial(0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})

	t.Run("negative input", func(t *testing.T) {
		_, err := Factorial(-1)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestFizzBuzz(t *testing.T) {
	t.Run("up to fifteen", func(t *testing.T) {
		want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
		got := FizzBuzz(15)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("non-positive max", func(t *testing.T) {
		got := FizzBuzz(0)
		if len(got) != 0 {
			t.Errorf("got %v, want empty slice", got)
		}
	})
}

func TestSumDigits(t *testing.T) {
	t.Run("positive number", func(t *testing.T) {
		if got := SumDigits(1234); got != 10 {
			t.Errorf("got %d, want 10", got)
		}
	})

	t.Run("negative number", func(t *testing.T) {
		if got := SumDigits(-909); got != 18 {
			t.Errorf("got %d, want 18", got)
		}
	})

	t.Run("zero", func(t *testing.T) {
		if got := SumDigits(0); got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})
}
