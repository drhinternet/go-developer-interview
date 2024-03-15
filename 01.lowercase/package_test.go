package lowercase_test

import (
	"testing"

	"lowercase"
)

func TestToLower(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"", ""},
		{"foo", "foo"},
		{"Hello World", "hello world"},
		{" X X X ", " x x x "},
	}

	for _, test := range tests {
		output := lowercase.ToLower(test.input)

		if output != test.expect {
			t.Errorf("expected %v, got %v", test.expect, output)
		}
	}
}
