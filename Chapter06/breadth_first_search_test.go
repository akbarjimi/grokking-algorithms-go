package Chapter06

import (
	"testing"
)

func TestBreadthFirstSearch_path(t *testing.T) {
	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"D"}
	graph["C"] = []string{"E"}
	graph["D"] = []string{"E"}
	graph["E"] = []string{}

	start := "A"
	needle := "E"
	expected := true

	t.Run("When there is a path", func(t *testing.T) {
		got := BreadthFirstSearch(graph, start, needle)
		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}

func TestBreadthFirstSearch_no_path(t *testing.T) {
	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"D"}
	graph["C"] = []string{"E"}
	graph["D"] = []string{"E"}
	graph["E"] = []string{}

	start := "A"
	needle := "Z"
	expected := false

	t.Run("When there is no path", func(t *testing.T) {
		got := BreadthFirstSearch(graph, start, needle)
		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}

func TestBreadthFirstSearch_just_one_path(t *testing.T) {
	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"D"}
	graph["C"] = []string{}
	graph["D"] = []string{"E"}
	graph["E"] = []string{}

	start := "A"
	needle := "E"
	expected := true

	t.Run("When there is no path", func(t *testing.T) {
		got := BreadthFirstSearch(graph, start, needle)
		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}
