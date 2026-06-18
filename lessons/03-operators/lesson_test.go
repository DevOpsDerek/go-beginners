package lesson03

import (
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {
	t.Run("addition", func(t *testing.T) {
		got, err := Calculate(3, "+", 4)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 7 {
			t.Errorf("got %v, want 7", got)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := Calculate(5, "/", 0)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if err.Error() != "division by zero" {
			t.Errorf("got %q, want division by zero", err.Error())
		}
	})

	t.Run("unknown operator", func(t *testing.T) {
		_, err := Calculate(1, "^", 2)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !strings.Contains(err.Error(), "unknown operator: ^") {
			t.Errorf("unexpected error message: %v", err)
		}
	})
}

func TestIsEven(t *testing.T) {
	t.Run("even number", func(t *testing.T) {
		if !IsEven(14) {
			t.Error("expected true for even number")
		}
	})

	t.Run("odd negative number", func(t *testing.T) {
		if IsEven(-3) {
			t.Error("expected false for odd number")
		}
	})
}

func TestIsPalindrome(t *testing.T) {
	t.Run("simple palindrome", func(t *testing.T) {
		if !IsPalindrome("level") {
			t.Error("expected level to be a palindrome")
		}
	})

	t.Run("unicode palindrome", func(t *testing.T) {
		if !IsPalindrome("été") {
			t.Error("expected unicode palindrome to return true")
		}
	})

	t.Run("not a palindrome", func(t *testing.T) {
		if IsPalindrome("gopher") {
			t.Error("expected false for non-palindrome")
		}
	})
}
