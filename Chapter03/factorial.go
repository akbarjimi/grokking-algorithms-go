package Chapter03

func Factorial(input int) int {
	if input == 0 || input == 1 {
		return 1
	}

	if input == 2 {
		return 2
	}

	return input * Factorial(input-1)
}
