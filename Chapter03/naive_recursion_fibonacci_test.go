package Chapter03

import "testing"

func TestNaiveRecursionFibonacci(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{name: "0 as input", input: 0, expected: 0},
		{name: "1 as input", input: 1, expected: 1},
		{name: "2 as input", input: 2, expected: 1},
		{name: "3 as input", input: 3, expected: 2},
		{name: "4 as input", input: 4, expected: 3},
		{name: "5 as input", input: 5, expected: 5},
		{name: "6 as input", input: 6, expected: 8},
		{name: "7 as input", input: 7, expected: 13},
		{name: "40 as input", input: 40, expected: 102334155},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := NaiveRecursionFibonacci(testCase.input)
			if result != testCase.expected {
				t.Errorf("Input: %v, Expected: %v, Actual: %v", testCase.input, testCase.expected, result)
			}
		})
	}
}
