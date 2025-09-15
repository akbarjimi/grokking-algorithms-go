package Chapter01

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name     string
		haystack []int
		needle   int
		want     int
	}{
		{
			name:     "Element in the middle",
			haystack: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			needle:   5,
			want:     4,
		},
		{
			name:     "Element at the start",
			haystack: []int{1, 2, 3, 4, 5},
			needle:   1,
			want:     0,
		},
		{
			name:     "Element at the end",
			haystack: []int{1, 2, 3, 4, 5},
			needle:   5,
			want:     4,
		},
		{
			name:     "Element not present",
			haystack: []int{10, 20, 30, 40, 50},
			needle:   25,
			want:     -1,
		},
		{
			name:     "Empty slice",
			haystack: []int{},
			needle:   1,
			want:     -1,
		},
		{
			name:     "Single element, found",
			haystack: []int{42},
			needle:   42,
			want:     0,
		},
		{
			name:     "Single element, not found",
			haystack: []int{42},
			needle:   10,
			want:     -1,
		},
		{
			name:     "Even number of elements, found",
			haystack: []int{1, 5, 10, 15, 20, 25},
			needle:   15,
			want:     3,
		},
		{
			name:     "Odd number of elements, found",
			haystack: []int{1, 5, 10, 15, 20},
			needle:   10,
			want:     2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BinarySearch(tc.needle, tc.haystack)

			if got != tc.want {
				t.Errorf("BinarySearch(%d, %v) = %d; want %d", tc.needle, tc.haystack, got, tc.want)
			}
		})
	}
}
