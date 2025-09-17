package Chapter03

import "testing"

func TestFactorial(t *testing.T) {
	testCases := []struct {
		name   string
		input  int
		output int
	}{
		{name: "0 as input", input: 0, output: 1},
		{name: "1 as input", input: 1, output: 1},
		{name: "2 as input", input: 2, output: 2},
		{name: "3 as input", input: 3, output: 6},
		{name: "4 as input", input: 4, output: 24},
		{name: "5 as input", input: 5, output: 120},
		{name: "20 as input", input: 20, output: 2432902008176640000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Factorial(tc.input)
			if got != tc.output {
				t.Errorf("For input %d: expected %d, got %d", tc.input, tc.output, got)
			}
		})
	}
}
