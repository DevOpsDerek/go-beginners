package lesson06

import (
	"fmt"
	"strings"
)

/*
Lesson 06: Slices & Maps
Learning objectives:
- Understand how slices differ from fixed-size arrays.
- Use append, len, and cap while growing data.
- Build and merge maps safely.
- Iterate with range over slices and maps.
Estimated time: 25 minutes
Prerequisites: Lessons 01-05.
*/

// SliceStats calculates basic statistics for a slice of numbers.
func SliceStats(nums []float64) (min, max, sum, avg float64, err error) {
	if len(nums) == 0 {
		return 0, 0, 0, 0, fmt.Errorf("empty slice")
	}

	min = nums[0]
	max = nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}

	avg = sum / float64(len(nums))
	return min, max, sum, avg, nil
}

// WordCount splits a string on whitespace and counts each word.
func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(s) {
		counts[word]++
	}

	return counts
}

// MergeMaps combines two maps into a new result map.
// Keys in override replace the same keys in base.
func MergeMaps(base, override map[string]string) map[string]string {
	merged := make(map[string]string, len(base)+len(override))
	for key, value := range base {
		merged[key] = value
	}
	for key, value := range override {
		merged[key] = value
	}

	return merged
}

// RunDemo prints slice and map examples.
func RunDemo() {
	min, max, sum, avg, _ := SliceStats([]float64{2, 4, 6})
	fmt.Println("Stats:", min, max, sum, avg)
	fmt.Println("WordCount:", WordCount("go is fun and go is fast"))
	fmt.Println("MergeMaps:", MergeMaps(map[string]string{"env": "dev"}, map[string]string{"env": "prod", "region": "us"}))
}

// ---- FOLLOW ALONG ----
// TODO: Call SliceStats with a longer slice.
// TODO: Try WordCount with repeated words and extra spaces.
// TODO: Add a helper that returns only the keys from a map.
