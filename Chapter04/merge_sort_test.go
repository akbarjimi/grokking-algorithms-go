package Chapter04

import (
	"reflect"
	"testing"
)

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMerge(t *testing.T) {
	testCases := []struct {
		name     string
		left     []int
		right    []int
		expected []int
	}{
		{name: "Empty slices", left: []int{}, right: []int{}, expected: []int{}},
		{name: "Left empty", left: []int{}, right: []int{1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "Right empty", left: []int{4, 5, 6}, right: []int{}, expected: []int{4, 5, 6}},
		{name: "Standard merge", left: []int{1, 5, 8}, right: []int{2, 6, 9}, expected: []int{1, 2, 5, 6, 8, 9}},
		{name: "Duplicate values", left: []int{2, 4, 6}, right: []int{2, 5, 7}, expected: []int{2, 2, 4, 5, 6, 7}},
		{name: "Uneven lengths", left: []int{1, 10, 11}, right: []int{2, 3}, expected: []int{1, 2, 3, 10, 11}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := merge(tc.left, tc.right)
			if !slicesEqual(got, tc.expected) {
				t.Errorf("merge(%v, %v): expected %v, got %v", tc.left, tc.right, tc.expected, got)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "Empty slice", input: []int{}, expected: []int{}},
		{name: "Single element", input: []int{5}, expected: []int{5}},
		{name: "Already sorted", input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{name: "Reverse sorted", input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{name: "Random small", input: []int{9, 2, 7, 4}, expected: []int{2, 4, 7, 9}},
		{name: "Random with duplicates", input: []int{8, 5, 1, 5, 9, 3, 8}, expected: []int{1, 3, 5, 5, 8, 8, 9}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := MergeSort(tc.input)

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("MergeSort(%v): expected %v, got %v", tc.input, tc.expected, got)
			}
		})
	}
}
