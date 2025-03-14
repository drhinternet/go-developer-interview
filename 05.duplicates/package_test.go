package duplicates_test

import (
	"duplicates"
	"slices"
	"testing"
)

func TestDuplicates(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{}},
		{[]int{1, 1, 1}, []int{1}},
		{[]int{5, 4, 5, 3, 2, 1}, []int{5}},
		{[]int{5, 4, 5, 3, 2, 1, 4}, []int{4, 5}},
	}

	for _, test := range tests {
		output := duplicates.FindDuplicates(test.input)

		if !slices.Equal(output, test.expect) {
			t.Errorf("expected %v, got %v", test.expect, output)
			continue
		}
	}
}
