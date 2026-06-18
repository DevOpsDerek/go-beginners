package lesson10

import (
	"math"
	"reflect"
	"testing"
)

func TestStudentGrade(t *testing.T) {
	t.Run("grade a", func(t *testing.T) {
		student := Student{Name: "Ava", Score: 95}
		if got := student.Grade(); got != "A" {
			t.Errorf("got %q, want A", got)
		}
	})

	t.Run("invalid score", func(t *testing.T) {
		student := Student{Name: "Ava", Score: 200}
		if got := student.Grade(); got != "invalid" {
			t.Errorf("got %q, want invalid", got)
		}
	})
}

func TestNewStudent(t *testing.T) {
	t.Run("valid student", func(t *testing.T) {
		got, err := NewStudent("Noah", 88)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Name != "Noah" || got.Score != 88 {
			t.Errorf("got %+v, want Name=Noah Score=88", got)
		}
	})

	t.Run("empty name", func(t *testing.T) {
		_, err := NewStudent("", 88)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("invalid score", func(t *testing.T) {
		_, err := NewStudent("Noah", 101)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClassAverage(t *testing.T) {
	t.Run("average of class", func(t *testing.T) {
		students := []Student{{Name: "A", Score: 80}, {Name: "B", Score: 90}, {Name: "C", Score: 100}}
		got, err := ClassAverage(students)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if math.Abs(got-90) > 1e-9 {
			t.Errorf("got %v, want 90", got)
		}
	})

	t.Run("empty class", func(t *testing.T) {
		_, err := ClassAverage(nil)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestTopStudents(t *testing.T) {
	t.Run("top two sorted descending", func(t *testing.T) {
		students := []Student{{Name: "A", Score: 75}, {Name: "B", Score: 98}, {Name: "C", Score: 88}}
		got := TopStudents(students, 2)
		want := []Student{{Name: "B", Score: 98}, {Name: "C", Score: 88}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("n larger than class", func(t *testing.T) {
		students := []Student{{Name: "A", Score: 75}}
		got := TopStudents(students, 5)
		if len(got) != 1 {
			t.Errorf("got %d students, want 1", len(got))
		}
	})

	t.Run("non-positive n", func(t *testing.T) {
		students := []Student{{Name: "A", Score: 75}}
		got := TopStudents(students, 0)
		if len(got) != 0 {
			t.Errorf("got %v, want empty slice", got)
		}
	})
}
