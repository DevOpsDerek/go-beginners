package lesson04

import "fmt"

/*
Lesson 04: Conditionals
Learning objectives:
- Write if, else if, and else branches.
- Use switch statements for readable branching.
- Understand validation before business logic.
- Practice returning helpful error messages.
Estimated time: 20 minutes
Prerequisites: Lessons 01-03.
*/

// GetSeason returns the season for a calendar month.
func GetSeason(month int) (string, error) {
	if month < 1 || month > 12 {
		return "", fmt.Errorf("invalid month: %d", month)
	}

	switch month {
	case 3, 4, 5:
		return "Spring", nil
	case 6, 7, 8:
		return "Summer", nil
	case 9, 10, 11:
		return "Autumn", nil
	default:
		return "Winter", nil
	}
}

// GetLetterGrade converts a numeric score into a letter grade.
func GetLetterGrade(score int) (string, error) {
	if score < 0 || score > 100 {
		return "", fmt.Errorf("score must be between 0 and 100")
	}

	switch {
	case score >= 90:
		return "A", nil
	case score >= 80:
		return "B", nil
	case score >= 70:
		return "C", nil
	case score >= 60:
		return "D", nil
	default:
		return "F", nil
	}
}

// Classify labels a number as negative, zero, or positive.
func Classify(n int) string {
	if n < 0 {
		return "negative"
	}
	if n == 0 {
		return "zero"
	}

	return "positive"
}

// RunDemo shows example branch results.
func RunDemo() {
	season, _ := GetSeason(7)
	grade, _ := GetLetterGrade(88)
	fmt.Println("Season for July:", season)
	fmt.Println("Grade for 88:", grade)
	fmt.Println("Classify -2:", Classify(-2))
}

// ---- FOLLOW ALONG ----
// TODO: Check which season month 12 belongs to.
// TODO: Try GetLetterGrade with scores on the boundary like 89 and 90.
// TODO: Add a new helper that classifies numbers as small, medium, or large.
