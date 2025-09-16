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

	var temp int
	loops := len(slice) - 1

	for i := 0; i < loops; i++ {
		minIndex, err := findMinimumIndexV2(slice, i)
		if err != nil {
			panic(err)
		}

		if i == minIndex {
			continue
		}

		temp = slice[i]
		slice[i] = slice[minIndex]
		slice[minIndex] = temp
	}

	return slice
}
