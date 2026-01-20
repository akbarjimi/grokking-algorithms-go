package chapter07

import "testing"

func TestDepthFirstSearchRecursive(t *testing.T) {
	tests := []struct {
		name     string
		graph    map[string][]string
		start    string
		needle   string
		expected bool
	}{
		{
			name: "single node",
			graph: map[string][]string{
				"E": {},
			},
			start:    "E",
			needle:   "E",
			expected: true,
		}, {
			name: "direct neighbor",
			graph: map[string][]string{
				"D": {"E"},
			},
			start:    "D",
			needle:   "E",
			expected: true,
		}, {
			name: "deep chain",
			graph: map[string][]string{
				"A": {"B"},
				"B": {"C"},
				"C": {"D"},
			},
			start:    "A",
			needle:   "D",
			expected: true,
		}, {
			name: "branching graph",
			graph: map[string][]string{
				"A": {"B", "C"},
				"B": {},
				"C": {"D"},
				"D": {},
			},
			start:    "A",
			needle:   "D",
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
			name: "unreachable",
			graph: map[string][]string{
				"A": {"B"},
			},
			start:    "A",
			needle:   "C",
			expected: false,
		}, {
			name: "start node not in graph",
			graph: map[string][]string{
				"A": {"B"},
			},
			start:    "Z",
			needle:   "B",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := DepthFirstSearchRecursive(test.graph, test.start, test.needle, make(map[string]bool))
			if got != test.expected {
				t.Fatalf(
					"DepthFirstSearchRecusrive(%q → %q) = %v, want %v",
					test.start,
					test.needle,
					got,
					test.expected,
				)
			}
		})
	}
}
