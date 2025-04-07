package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		num      int
		expected int
	}{
		{"case 2", 2, 2},
		{"case 5", 5, 120},
		{"case -1", -1, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Factorial(tc.num)
			expectedResult := tc.expected

			if result != expectedResult {
				t.Errorf("Factorial(%d) = %d is wrong, correct is %d", tc.num, result, expectedResult)
			}
		})
	}
}
