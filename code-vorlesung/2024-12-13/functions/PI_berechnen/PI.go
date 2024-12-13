package piberechnen

import "math"

func PI(n int) float64 {
	var result float64 = 0
	for i := 0; i < n; i++ {
		p := float64(i)
		result = +math.Pow(-1, p) / (2*p + 1)
	}
	return result
}
