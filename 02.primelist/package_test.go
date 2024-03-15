package primelist_test

import (
	"slices"
	"testing"

	"primelist"
)

func TestPrimes(t *testing.T) {
	tests := []struct {
		input  int
		expect []int
	}{
		{0, []int{}},
		{1, []int{2}},
		{2, []int{2, 3}},
		{10, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	}

	for _, test := range tests {
		output := primelist.Primes(test.input)

		if !slices.Equal(output, test.expect) {
			t.Errorf("expected %d primes %v, got %d primes %v", len(test.expect), test.expect, len(output), output)
		}
	}
}
