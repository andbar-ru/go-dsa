package dsa

func factorial(n int) int {
	fact := 1
	for v := 1; v <= n; v++ {
		fact *= v
	}
	return fact
}
