package lesson10

import (
	"fmt"
	"sort"
)

/*
Lesson 10: Capstone - Structs & Interfaces
Learning objectives:
- Define and validate structs.
- Attach methods to struct types.
- Satisfy interfaces with methods.
- Sort and summarize collections of structs.
Estimated time: 30 minutes
Prerequisites: Lessons 01-09.
*/

// Student stores a learner name and score.
type Student struct {
	Name  string
	Score int
}

// Grader is implemented by any type with a Grade method.
type Grader interface {
	Grade() string
}

// Grade converts a numeric score to a letter grade.
func (s Student) Grade() string {
	switch {
	case s.Score < 0 || s.Score > 100:
		return "invalid"
	case s.Score >= 90:
		return "A"
	case s.Score >= 80:
		return "B"
	case s.Score >= 70:
		return "C"
	case s.Score >= 60:
		return "D"
	default:
		return "F"
	}
}

// NewStudent validates the input before returning a Student value.
func NewStudent(name string, score int) (Student, error) {
	if name == "" {
		return Student{}, fmt.Errorf("name cannot be empty")
	}
	if score < 0 || score > 100 {
		return Student{}, fmt.Errorf("score must be between 0 and 100")
	}

	return Student{Name: name, Score: score}, nil
}

// ClassAverage calculates the average score for a class.
func ClassAverage(students []Student) (float64, error) {
	if len(students) == 0 {
		return 0, fmt.Errorf("empty class")
	}

	total := 0
	for _, student := range students {
		total += student.Score
	}

	return float64(total) / float64(len(students)), nil
}

// TopStudents returns a new slice sorted by descending score and trimmed to n items.
func TopStudents(students []Student, n int) []Student {
	if n <= 0 || len(students) == 0 {
		return []Student{}
	}

	result := append([]Student(nil), students...)
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Score > result[j].Score
	})

	if n > len(result) {
		n = len(result)
	}

	return result[:n]
}

// RunDemo prints sample struct and interface usage.
func RunDemo() {
	student, _ := NewStudent("Avery", 92)
	var grader Grader = student
	fmt.Println("Student grade:", grader.Grade())
	average, _ := ClassAverage([]Student{{Name: "A", Score: 80}, {Name: "B", Score: 90}})
	fmt.Println("Class average:", average)
	fmt.Println("Top student:", TopStudents([]Student{{Name: "A", Score: 80}, {Name: "B", Score: 90}}, 1))
}

// ---- FOLLOW ALONG ----
// TODO: Create three students and calculate the class average.
// TODO: Check which grade a score of 59 receives.
// TODO: Add another method to Student, such as Passed() bool.
