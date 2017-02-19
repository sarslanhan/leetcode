package hammingdistance

import (
	"testing"
)

type TestCase struct {
	x        int
	y        int
	distance int
}

var testCases = [...]TestCase{
	TestCase{1, 4, 2},
	TestCase{0, 1, 1},
}

func TestHammingDistance(t *testing.T) {
	for _, tc := range testCases {
		result := hammingDistance(tc.x, tc.y)
		if result != tc.distance {
			t.Fatalf("Expected %d, got %d for %d - %d", tc.distance, result, tc.x, tc.y)
		}
	}
}
