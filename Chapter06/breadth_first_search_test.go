package chapter06

import (
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	tests := []struct {
		name     string
		graph    map[string][]string
		start    string
		needle   string
		expected bool
	}{
		{
			name: "path exists with multiple branches",
			graph: map[string][]string{
				"A": {"B", "C"},
				"B": {"D"},
				"C": {"E"},
				"D": {"E"},
				"E": {},
			},
			start:    "A",
			needle:   "E",
			expected: true,
		}, {
			name: "no path exists",
			graph: map[string][]string{
				"A": {"B", "C"},
				"B": {"D"},
				"C": {"E"},
				"D": {"E"},
				"E": {},
			},
			start:    "A",
			needle:   "Z",
			expected: false,
		}, {
			name: "single narrow path",
			graph: map[string][]string{
				"A": {"B", "C"},
				"B": {"D"},
				"C": {},
				"D": {"E"},
				"E": {},
			},
			start:    "A",
			needle:   "E",
			expected: true,
		}, {
			name: "start equals needle",
			graph: map[string][]string{
				"A": {"B"},
				"B": {"C"},
			},
			start:    "A",
			needle:   "A",
			expected: true,
		}, {
			name: "graph with cycle",
			graph: map[string][]string{
				"A": {"B"},
				"B": {"C"},
				"C": {"A"},
			},
			start:    "A",
			needle:   "C",
			expected: true,
		}, {
			name:     "empty graph",
			graph:    map[string][]string{},
			start:    "A",
			needle:   "B",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := BreadthFirstSearch(test.graph, test.start, test.needle)
			if got != test.expected {
				t.Fatalf(
					"BreadthFirstSearch(%q → %q) = %v, want %v",
					test.start,
					test.needle,
					got,
					test.expected,
				)
			}
		})
	}
}
