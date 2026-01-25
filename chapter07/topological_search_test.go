package chapter07

import "testing"

func TestTopologicalSortSimple(t *testing.T) {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D"},
		"C": {"D"},
		"D": {},
	}

	res := TopologicalSort(graph)

	// valid orders:
	// A B C D
	// A C B D

	if len(res) != 4 {
		t.Fatalf("expected 4 nodes, got %d", len(res))
	}

	if res[0] != "A" {
		t.Fatalf("A should be first, got %s", res[0])
	}

	if res[3] != "D" {
		t.Fatalf("D should be last, got %s", res[3])
	}
}

func TestSingleNode(t *testing.T) {
	graph := map[string][]string{
		"A": {},
	}

	res := TopologicalSort(graph)

	if len(res) != 1 || res[0] != "A" {
		t.Fatal("single node graph failed")
	}
}
