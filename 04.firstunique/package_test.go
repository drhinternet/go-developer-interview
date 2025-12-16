package firstunique_test

import (
	"firstunique"
	"testing"
)

func TestFirstUnique(t *testing.T) {
	tests := []struct {
		input  string
		expect rune
	}{
		{"", 0},
		{"x", 'x'},
		{"fff", 0},
		{"foo", 'f'},
		{"oof", 'f'},
		{"food", 'f'},
		{"doof", 'd'},
		{"xyxzx", 'y'},
	}

	for _, test := range tests {
		output := firstunique.FirstUniqueRune(test.input)

		if output != test.expect {
			t.Errorf("expected %q, got %q", test.expect, output)
			continue
		}
	}
}
