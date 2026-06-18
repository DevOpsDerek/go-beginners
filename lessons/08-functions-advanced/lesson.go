package lesson08

import "fmt"

/*
Lesson 08: Functions Advanced
Learning objectives:
- Build closures that remember local state.
- Return functions from other functions.
- Understand defer and its last-in, first-out order.
- Use anonymous functions for short, focused logic.
Estimated time: 25 minutes
Prerequisites: Lessons 01-07.
*/

// MakeCounter returns a closure that increments a private counter.
func MakeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// MakeAdder returns a closure that adds n to its argument.
func MakeAdder(n int) func(int) int {
	return func(value int) int {
		return value + n
	}
}

// RunWithDefer demonstrates that deferred calls run in reverse order.
func RunWithDefer() (steps []string) {
	steps = append(steps, "start")
	defer func() {
		steps = append(steps, "deferred 1")
	}()
	defer func() {
		steps = append(steps, "deferred 2")
	}()
	steps = append(steps, "end")

	return steps
}

// RunDemo prints closure and defer examples.
func RunDemo() {
	counter := MakeCounter()
	addFive := MakeAdder(5)
	fmt.Println("Counter calls:", counter(), counter())
	fmt.Println("Add five:", addFive(10))
	fmt.Println("Defer order:", RunWithDefer())
}

// ---- FOLLOW ALONG ----
// TODO: Create two counters and notice that each keeps separate state.
// TODO: Make an adder that adds 10.
// TODO: Add one more deferred function and predict its position in the result.
