package dsa

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func checkPrime(n int) bool {
	if n == 1 || n == 0 {
		return false
	}
	for i := 2; i < n/2; i++ {
		if n%2 == 0 {
			return false
		}
	}
	return true
}

func getPrime(n int) int {
	if n%2 == 0 {
		n++
	}
	for !checkPrime(n) {
		n += 2
	}
	return n
}
