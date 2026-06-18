package lesson09

import (
	"errors"
	"testing"
)

func TestSafeDivide(t *testing.T) {
	t.Run("successful division", func(t *testing.T) {
		got, err := SafeDivide(12, 3)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 4 {
			t.Errorf("got %v, want 4", got)
		}
	})

	t.Run("division by zero wraps sentinel", func(t *testing.T) {
		_, err := SafeDivide(12, 0)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, errDivisionByZero) {
			t.Errorf("expected wrapped division-by-zero error, got %v", err)
		}
	})
}

func TestValidateAge(t *testing.T) {
	t.Run("valid age", func(t *testing.T) {
		if err := ValidateAge(30); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("invalid age returns validation error", func(t *testing.T) {
		err := ValidateAge(151)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		var validationErr *ValidationError
		if !errors.As(err, &validationErr) {
			t.Fatalf("expected ValidationError, got %T", err)
		}
		if validationErr.Field != "age" {
			t.Errorf("got field %q, want age", validationErr.Field)
		}
	})
}

func TestSafeRun(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		if err := SafeRun(func() {}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("panic becomes error", func(t *testing.T) {
		err := SafeRun(func() { panic("boom") })
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if err.Error() != "panic: boom" {
			t.Errorf("got %q, want %q", err.Error(), "panic: boom")
		}
	})
}
