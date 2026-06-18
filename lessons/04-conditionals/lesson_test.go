package lesson04

import "testing"

func TestGetSeason(t *testing.T) {
	t.Run("spring month", func(t *testing.T) {
		got, err := GetSeason(4)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "Spring" {
			t.Errorf("got %q, want Spring", got)
		}
	})

	t.Run("winter month", func(t *testing.T) {
		got, err := GetSeason(1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "Winter" {
			t.Errorf("got %q, want Winter", got)
		}
	})

	t.Run("invalid month", func(t *testing.T) {
		_, err := GetSeason(13)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestGetLetterGrade(t *testing.T) {
	t.Run("grade a", func(t *testing.T) {
		got, err := GetLetterGrade(95)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "A" {
			t.Errorf("got %q, want A", got)
		}
	})

	t.Run("grade f", func(t *testing.T) {
		got, err := GetLetterGrade(40)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "F" {
			t.Errorf("got %q, want F", got)
		}
	})

	t.Run("out of range", func(t *testing.T) {
		_, err := GetLetterGrade(101)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClassify(t *testing.T) {
	t.Run("negative", func(t *testing.T) {
		if got := Classify(-1); got != "negative" {
			t.Errorf("got %q, want negative", got)
		}
	})

	t.Run("zero", func(t *testing.T) {
		if got := Classify(0); got != "zero" {
			t.Errorf("got %q, want zero", got)
		}
	})

	t.Run("positive", func(t *testing.T) {
		if got := Classify(1); got != "positive" {
			t.Errorf("got %q, want positive", got)
		}
	})
}
