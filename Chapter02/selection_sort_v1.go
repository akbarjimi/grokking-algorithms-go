package Chapter02

import (
	"errors"
)

func findMinimumIndex(slice []uint8) (uint8, error) {

	if len(slice) == 0 {
		return 0, errors.New("slice can not be empty")
	}

	var minimumIndex uint8 = 0

	for i := 0; i < len(slice); i++ {
		if slice[i] < slice[minimumIndex] {
			minimumIndex = uint8(i)
		}
	}

	return minimumIndex, nil
}

func popMinimumIndex(index uint8, slice []uint8) (uint8, []uint8, error) {

	if index < 0 || index >= uint8(len(slice)) {
		return 0, nil, errors.New("index out of bounds")
	}

	var newSlice []uint8
	var poppedValue uint8

	poppedValue = slice[index]
	newSlice = append(slice[:index], slice[index+1:]...)
	return poppedValue, newSlice, nil
}

func SelectionSort(slice []uint8) []uint8 {
	if len(slice) < 2 {
		return slice
	}

	var newSlice []uint8
	var minValue uint8
	loops := len(slice)

	for i := 0; i < loops; i++ {
		minIndex, err := findMinimumIndex(slice)
		if err != nil {
			panic(err)
		}

		minValue, slice, err = popMinimumIndex(minIndex, slice)
		if err != nil {
			panic(err)
		}

		newSlice = append(newSlice, minValue)
	}

	return newSlice
}
