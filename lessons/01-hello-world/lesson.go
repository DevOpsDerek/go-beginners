package lesson01

import "fmt"

/*
Lesson 01: Hello World & Packages
Learning objectives:
- Understand package declarations and imports.
- Use fmt.Println and fmt.Sprintf for formatted output.
- Recognize exported versus unexported identifiers.
- Review single-line and multi-line comments.
Estimated time: 15 minutes
Prerequisites: A working Go installation and basic terminal familiarity.
*/

// Greet returns a friendly greeting for the provided name.
//
// Because the function name starts with a capital letter, it is exported
// and can be used by code in other packages.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// FullName joins a first and last name with a single space.
func FullName(first, last string) string {
	return first + " " + last
}

// RunDemo prints a few examples that a beginner can read alongside the code.
// It is not a main function, so this package stays importable and testable.
func RunDemo() {
	fmt.Println(Greet("Go learner"))
	fmt.Println(FullName("Ada", "Lovelace"))
}

// ---- FOLLOW ALONG ----
// TODO: Call Greet with your own name.
// TODO: Create a variable for your first and last name and pass both to FullName.
// TODO: Add a new exported function that returns a goodbye message.
