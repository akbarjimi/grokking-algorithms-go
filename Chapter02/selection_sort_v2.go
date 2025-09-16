package Chapter02

import (
	"errors"
)

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
