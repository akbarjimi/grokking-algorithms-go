package Chapter02

import (
	"reflect"
	"testing"
)

func TestV1SelectionSort(t *testing.T) {
	testCases := []struct {
		name     string
		haystack []uint8
		want     []uint8
	}{
		{
			name:     "Sort these",
			haystack: []uint8{10, 8, 6, 4, 3, 2, 1, 5, 7, 9},
			want:     []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Basic case",
			haystack: []uint8{64, 25, 12, 22, 11},
			want:     []uint8{11, 12, 22, 25, 64},
		},
		{
			name:     "Already sorted slice",
			haystack: []uint8{1, 2, 3, 4, 5},
			want:     []uint8{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted slice",
			haystack: []uint8{5, 4, 3, 2, 1},
			want:     []uint8{1, 2, 3, 4, 5},
		},
		{
			name:     "Slice with duplicates",
			haystack: []uint8{5, 2, 5, 1, 2, 1},
			want:     []uint8{1, 1, 2, 2, 5, 5},
		},
		{
			name:     "Single element slice",
			haystack: []uint8{42},
			want:     []uint8{42},
		},
		{
			name:     "Empty slice",
			haystack: []uint8{},
			want:     []uint8{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := SelectionSort(testCase.haystack)

			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("For haystack %v, want %v, but got %v", testCase.haystack, testCase.want, got)
			}
		})
	}
}
