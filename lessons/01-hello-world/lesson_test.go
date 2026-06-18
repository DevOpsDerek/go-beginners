package lesson01

import "testing"

func TestGreet(t *testing.T) {
	t.Run("name provided", func(t *testing.T) {
		got := Greet("Derek")
		want := "Hello, Derek!"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("empty name", func(t *testing.T) {
		got := Greet("")
		want := "Hello, !"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestFullName(t *testing.T) {
	t.Run("first and last name", func(t *testing.T) {
		got := FullName("Ada", "Lovelace")
		want := "Ada Lovelace"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("empty last name", func(t *testing.T) {
		got := FullName("Prince", "")
		want := "Prince "
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
