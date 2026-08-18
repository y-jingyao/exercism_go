package armstrongnumbers

func intPow(n, m int) int {
	re := 1
	for i := 1; i <= m; i++ {
		re *= n
	}
	return re
}

func IsNumber(n int) bool {
	//panic("Please implement the IsNumber function")
	sum := 0
	m := 0
	tmp := n
	for tmp > 0 {
		tmp /= 10
		m++
	}
	tmp = n
	for tmp > 0 {
		sum += intPow(tmp%10, m)
		tmp /= 10
	}
	return sum == n
}
