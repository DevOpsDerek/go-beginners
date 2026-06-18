package lesson07

import (
	"fmt"
	"math"
)

/*
Lesson 07: Functions Basics
Learning objectives:
- Read and write function signatures.
- Return multiple values from one function.
- Use variadic parameters for flexible input.
- Pass functions as arguments.
Estimated time: 25 minutes
Prerequisites: Lessons 01-06.
*/

// CircleArea returns the area of a circle.
func CircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

// MinMax returns the smallest and largest values from nums.
func MinMax(nums ...int) (min, max int, err error) {
	if len(nums) == 0 {
		return 0, 0, fmt.Errorf("at least one number is required")
	}

	min = nums[0]
	max = nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max, nil
}

// Apply calls fn for each element and returns a new slice of results.
// If fn is nil, Apply returns a copy of nums unchanged.
func Apply(nums []int, fn func(int) int) []int {
	if nums == nil {
		return nil
	}

	results := make([]int, len(nums))
	for i, num := range nums {
		if fn == nil {
			results[i] = num
			continue
		}
		results[i] = fn(num)
	}

	return results
}

// RunDemo shows a few function-based examples.
func RunDemo() {
	fmt.Println("Circle area:", CircleArea(2))
	min, max, _ := MinMax(9, 4, 7, 1)
	fmt.Println("MinMax:", min, max)
	fmt.Println("Apply square:", Apply([]int{1, 2, 3}, func(n int) int { return n * n }))
}

// ---- FOLLOW ALONG ----
// TODO: Call MinMax with just one number.
// TODO: Pass a function to Apply that doubles each number.
// TODO: Add a new variadic Sum helper.
