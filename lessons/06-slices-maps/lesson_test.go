package lesson06

import (
	"reflect"
	"testing"
)

func TestSliceStats(t *testing.T) {
	t.Run("basic stats", func(t *testing.T) {
		min, max, sum, avg, err := SliceStats([]float64{2, 4, 6, 8})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if min != 2 || max != 8 || sum != 20 || avg != 5 {
			t.Errorf("got (%v, %v, %v, %v), want (2, 8, 20, 5)", min, max, sum, avg)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		_, _, _, _, err := SliceStats(nil)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if err.Error() != "empty slice" {
			t.Errorf("got %q, want empty slice", err.Error())
		}
	})
}

func TestWordCount(t *testing.T) {
	t.Run("counts repeated words", func(t *testing.T) {
		got := WordCount("go go gopher")
		want := map[string]int{"go": 2, "gopher": 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("empty string", func(t *testing.T) {
		got := WordCount("")
		if len(got) != 0 {
			t.Errorf("got %v, want empty map", got)
		}
	})
}

func TestMergeMaps(t *testing.T) {
	t.Run("override wins", func(t *testing.T) {
		base := map[string]string{"env": "dev", "region": "west"}
		override := map[string]string{"env": "prod"}
		got := MergeMaps(base, override)
		want := map[string]string{"env": "prod", "region": "west"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("does not mutate input", func(t *testing.T) {
		base := map[string]string{"name": "course"}
		got := MergeMaps(base, nil)
		got["name"] = "changed"
		if base["name"] != "course" {
			t.Errorf("base map was mutated: %v", base)
		}
	})
}
