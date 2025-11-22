package Chapter04

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty slice", []int{}, []int{}},
		{"Single element", []int{5}, []int{5}},
		{"Already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"Random small", []int{9, 2, 7, 4}, []int{2, 4, 7, 9}},
		{"Random with duplicates", []int{8, 5, 1, 5, 9, 3, 8}, []int{1, 3, 5, 5, 8, 8, 9}},
		{"All same", []int{7, 7, 7, 7}, []int{7, 7, 7, 7}},
		{"Negatives", []int{-3, 10, -1, 0, 5}, []int{-3, -1, 0, 5, 10}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := QuickSort(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("QuickSort(%v): expected %v, got %v", tc.input, tc.expected, got)
			}
		})
	}
}

func TestQuickSort_LargeRandom(t *testing.T) {
	n := 100_000
	input := make([]int, n)
	for i := range input {
		input[i] = rand.Intn(1_000_000)
	}

	sorted := QuickSort(input)

	if sorted[0] > sorted[1] || sorted[n-2] > sorted[n-1] {
		t.Errorf("QuickSort failed on large random array")
	}
}

func TestQuickSort_ManyDuplicates(t *testing.T) {
	n := 100_000
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = i % 5
	}
	sorted := QuickSort(input)
	for i := 1; i < n; i++ {
		if sorted[i] < sorted[i-1] {
			t.Errorf("QuickSort failed with many duplicates")
		}
	}
}

func TestQuickSort_SortedReverse(t *testing.T) {
	n := 100_000
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = i
	}
	_ = QuickSort(input)

	for i := 0; i < n; i++ {
		input[i] = n - i
	}
	_ = QuickSort(input)
}
