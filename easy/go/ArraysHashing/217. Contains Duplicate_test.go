// 217_ContainsDuplicate_test.go
package ArraysHashing

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"No duplicates", []int{1, 2, 3, 4, 5}, false},
		{"One duplicate", []int{1, 2, 3, 4, 1}, true},
		{"Multiple duplicates", []int{1, 1, 2, 3, 3, 4, 5}, true},
		{"Empty slice", []int{}, false},
		{"Single element", []int{1}, false},
	}

	// Iterate through test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := containsDuplicate(tc.input)
			if result != tc.expected {
				t.Errorf("For input %v, expected %v but got %v", tc.input, tc.expected, result)
			}
		})
	}
}
