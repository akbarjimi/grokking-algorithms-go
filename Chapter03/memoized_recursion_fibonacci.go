package Chapter03

var results map[int]int = make(map[int]int)

func MemoizedRecursionFibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if value, ok := results[n]; ok {
		return value
	}
	results[n] = MemoizedRecursionFibonacci(n-1) + MemoizedRecursionFibonacci(n-2)
	return results[n]
}
