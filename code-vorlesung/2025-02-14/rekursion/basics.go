package rekursion

import "fmt"

func CountDown(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	CountDown(n - 1)
}

func CountUp(n int) {
	if n <= 0 {
		return
	}
	CountUp(n - 1)
	fmt.Println(n)

}

func Subtract(n int) int {
	if n <= 0 {
		return 0
	}
	return n - Subtract(n-1)
}

// Ergebnis der Basis b zur Potenz p
func Potenz(b, p float64) float64 {
	if p == 0 {
		return 1
	}
	if p < 0 {
		return 1 / Potenz(b, -p)
	}

	return b * Potenz(b, p-1)
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Fibonacci(n int) int {
	if n == 1|2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
