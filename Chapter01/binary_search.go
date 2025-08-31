package Chapter01

func BinarySearch[T ~int | ~int8 | ~int16 | ~int32 | ~int64](needle T, haystack []T) int {
	var low = 0
	var high = len(haystack) - 1

	for low <= high {
		middle := low + (high-low)/2
		if needle == haystack[middle] {
			return middle
		} else if needle < haystack[middle] {
			high = middle - 1
		} else if needle > haystack[middle] {
			low = middle + 1
		}
	}

	return -1
}
