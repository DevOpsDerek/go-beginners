package lesson07

import (
	"math"
	"reflect"
	"testing"
)

func TestCircleArea(t *testing.T) {
	t.Run("radius two", func(t *testing.T) {
		got := CircleArea(2)
		want := math.Pi * 4
		if math.Abs(got-want) > 1e-9 {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("zero radius", func(t *testing.T) {
		if got := CircleArea(0); got != 0 {
			t.Errorf("got %v, want 0", got)
		}
	})
}

func TestMinMax(t *testing.T) {
	t.Run("multiple numbers", func(t *testing.T) {
		min, max, err := MinMax(7, 2, 9, -1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if min != -1 || max != 9 {
			t.Errorf("got (%d, %d), want (-1, 9)", min, max)
		}
	})

	t.Run("no numbers", func(t *testing.T) {
		_, _, err := MinMax()
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestApply(t *testing.T) {
	t.Run("transform slice", func(t *testing.T) {
		got := Apply([]int{1, 2, 3}, func(n int) int { return n * 2 })
		want := []int{2, 4, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("nil function returns copy", func(t *testing.T) {
		input := []int{4, 5}
		got := Apply(input, nil)
		if !reflect.DeepEqual(got, input) {
			t.Errorf("got %v, want %v", got, input)
		}
		got[0] = 99
		if input[0] != 4 {
			t.Errorf("input slice was mutated: %v", input)
		}
	})
}
