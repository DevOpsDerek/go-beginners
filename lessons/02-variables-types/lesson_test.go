package lesson02

import (
	"math"
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	t.Run("freezing point", func(t *testing.T) {
		got := CelsiusToFahrenheit(0)
		if math.Abs(got-32) > 1e-9 {
			t.Errorf("got %v, want 32", got)
		}
	})

	t.Run("negative temperature", func(t *testing.T) {
		got := CelsiusToFahrenheit(-40)
		if math.Abs(got-(-40)) > 1e-9 {
			t.Errorf("got %v, want -40", got)
		}
	})
}

func TestTypeName(t *testing.T) {
	t.Run("string type", func(t *testing.T) {
		if got := TypeName("go"); got != "string" {
			t.Errorf("got %q, want string", got)
		}
	})

	t.Run("slice type", func(t *testing.T) {
		if got := TypeName([]int{1, 2, 3}); got != "[]int" {
			t.Errorf("got %q, want []int", got)
		}
	})
}

func TestAbsoluteValue(t *testing.T) {
	t.Run("negative number", func(t *testing.T) {
		if got := AbsoluteValue(-9); got != 9 {
			t.Errorf("got %d, want 9", got)
		}
	})

	t.Run("zero", func(t *testing.T) {
		if got := AbsoluteValue(0); got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})
}
