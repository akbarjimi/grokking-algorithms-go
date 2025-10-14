package Chapter04

func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	mid := len(slice) / 2

	left := MergeSort(slice[:mid])
	right := MergeSort(slice[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	// This prevents the final slice from being padded with zeros (Go's default initialization).
	results := make([]int, 0, len(left)+len(right))

	leftIndex, rightIndex := 0, 0

	for rightIndex < len(right) && leftIndex < len(left) {
		if left[leftIndex] <= right[rightIndex] {
			results = append(results, left[leftIndex])
			leftIndex++
		} else {
			results = append(results, right[rightIndex])
			rightIndex++
		}
	}

	// Appending an empty slice is a no-op
	results = append(results, right[rightIndex:]...)
	results = append(results, left[leftIndex:]...)
	return results
}
