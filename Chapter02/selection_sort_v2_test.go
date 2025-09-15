package Chapter02

import (
	"reflect"
	"testing"
)

func TestV2SelectionSort(t *testing.T) {
	testCases := []struct {
		name     string
		haystack []int
		want     []int
	}{
		{
			name:     "Sort these",
			haystack: []int{10, 8, 6, 4, 3, 2, 1, 5, 7, 9},
			want:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Basic case",
			haystack: []int{64, 25, 12, 22, 11},
			want:     []int{11, 12, 22, 25, 64},
		},
		{
			name:     "Already sorted slice",
			haystack: []int{1, 2, 3, 4, 5},
			want:     []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted slice",
			haystack: []int{5, 4, 3, 2, 1},
			want:     []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Slice with duplicates",
			haystack: []int{5, 2, 5, 1, 2, 1},
			want:     []int{1, 1, 2, 2, 5, 5},
		},
		{
			name:     "Single element slice",
			haystack: []int{42},
			want:     []int{42},
		},
		{
			name:     "Empty slice",
			haystack: []int{},
			want:     []int{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := SelectionSortV2(testCase.haystack)

			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("For haystack %v, want %v, but got %v", testCase.haystack, testCase.want, got)
			}
		})
	}
}
