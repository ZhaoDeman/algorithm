package math

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 反转字符串
func reverse(str string) string {
	b := []byte(str)
	n := len(b) - 1
	for i := 0; i < len(b)/2; i++ {
		b[i], b[n] = b[n], b[i]
		n--
	}
	return string(b)
}

// 快速判断是不是2的幂
func isPowTwo(n int) bool {
	return n&(n-1) == 0
}
