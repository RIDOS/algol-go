package fibonacci_number

// O(2^n)
func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}

	return fib(n-1) + fib(n-2)
}

func fibV2(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	var a, b int = 1, 1

	for i := 2; i < n; i++ {
		a, b = b, a+b
	}

	return b
}
