package Chapter02

import (
	"errors"
)

/*
* 1. In-Place Sorting: Think about how you can achieve the same result by just swapping elements within the original slice.
* 2. A "Sorted" and "Unsorted" Partition: Selection sort works by mentally dividing the slice into two parts: a sorted sub-slice at the beginning and an unsorted sub-slice at the end. Your outer loop should iterate through the unsorted portion, find the smallest element, and then move it to the end of the sorted portion.
* 3. One Function, One Purpose: Can you perform the sorting logic within the main function without needing helper functions that create new slices? What about a single loop with an inner loop?
 */
func findMinimumIndexV2(slice []int, startIndex int) (int, error) {

	if len(slice) == 0 {
		return 0, errors.New("slice can not be empty")
	}

	minimumIndex := startIndex

	for i := startIndex; i < len(slice); i++ {
		if slice[i] < slice[minimumIndex] {
			minimumIndex = i
		}
	}

	return minimumIndex, nil
}

func popMinimumIndexV2(index int, slice []int) (int, []int, error) {

	if index < 0 || index >= int(len(slice)) {
		return 0, nil, errors.New("index out of bounds")
	}

	var newSlice []int
	var poppedValue int

	poppedValue = slice[index]
	newSlice = append(slice[:index], slice[index+1:]...)
	return poppedValue, newSlice, nil
}

func SelectionSortV2(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	var newSlice []int
	var minValue int
	loops := len(slice)

	for i := 0; i < loops; i++ {
		minIndex, err := findMinimumIndexV2(slice, 0)
		if err != nil {
			panic(err)
		}

		minValue, slice, err = popMinimumIndexV2(minIndex, slice)
		if err != nil {
			panic(err)
		}

		newSlice = append(newSlice, minValue)
	}

	return newSlice
}
