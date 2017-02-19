package numbercomplement

import (
	"testing"
)

type TestCase struct {
	input  int
	output int
}

var testCases = [...]TestCase{
	TestCase{5, 2},
	TestCase{1, 0},
}

func TestFindComplement(t *testing.T) {
	for _, tc := range testCases {
		got := findComplement(tc.input)

		if got != tc.output {
			t.Errorf("Expected %d, got %d", tc.output, got)
		}
	}
}
