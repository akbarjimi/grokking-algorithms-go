package Chapter03

func NaiveRecursionFibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return NaiveRecursionFibonacci(n-1) + NaiveRecursionFibonacci(n-2)
}
