package lesson08

import (
	"reflect"
	"testing"
)

func TestMakeCounter(t *testing.T) {
	t.Run("counter increments independently", func(t *testing.T) {
		counter := MakeCounter()
		if got := counter(); got != 1 {
			t.Fatalf("first call got %d, want 1", got)
		}
		if got := counter(); got != 2 {
			t.Fatalf("second call got %d, want 2", got)
		}
	})

	t.Run("separate closures keep separate state", func(t *testing.T) {
		one := MakeCounter()
		two := MakeCounter()
		_ = one()
		if got := two(); got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})
}

func TestMakeAdder(t *testing.T) {
	t.Run("adds base value", func(t *testing.T) {
		addThree := MakeAdder(3)
		if got := addThree(7); got != 10 {
			t.Errorf("got %d, want 10", got)
		}
	})

	t.Run("negative adjustment", func(t *testing.T) {
		subtractTwo := MakeAdder(-2)
		if got := subtractTwo(5); got != 3 {
			t.Errorf("got %d, want 3", got)
		}
	})
}

func TestRunWithDefer(t *testing.T) {
	t.Run("defer order is reverse", func(t *testing.T) {
		got := RunWithDefer()
		want := []string{"start", "end", "deferred 2", "deferred 1"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
