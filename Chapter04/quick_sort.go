package Chapter04

import (
	"math/rand"
)

/*
Suggested QuickSort improvement chain:

1. Convert to in-place quicksort
   - Reduces memory allocations drastically
   - Operates on the same slice without creating new slices

2. Use contiguous memory
   - Improves CPU cache locality
   - All operations happen in the same slice

3. Median-of-three pivot selection
   - Picks a better pivot (first, middle, last elements)
   - Reduces recursion depth
   - Helps avoid worst-case patterns

4. Switch to insertion sort for small partitions
   - Avoids recursion overhead on small arrays
   - Very fast for slices of size ~10-20

5. Three-way partitioning (Dutch National Flag)
   - Efficiently handles duplicates
   - Reduces recursion depth for repeated elements

6. Introsort / recursion depth monitoring
   - If recursion depth exceeds 2*log2(n), switch to heapsort
   - Guarantees worst-case O(n log n) time

7. Optional CPU-level micro-optimizations
   - Branchless comparisons
   - Loop unrolling for small partitions

8. Benchmarks & stress tests
   - Test large arrays, many duplicates, already sorted/reverse arrays
   - Measure memory and execution time
*/

func QuickSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	randomIndex := rand.Intn(len(slice))

	pick := slice[randomIndex]

	var left []int
	var right []int
	var middle []int

	for _, v := range slice {
		if v < pick {
			left = append(left, v)
		}
		if v > pick {
			right = append(right, v)
		}
		if v == pick {
			middle = append(middle, v)
		}
	}

	var result []int

	result = append(result, QuickSort(left)...)
	result = append(result, middle...)
	result = append(result, QuickSort(right)...)
	return result
}
