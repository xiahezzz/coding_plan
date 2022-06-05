package coding

func CalFactorial(n int) int {
	temp := 1
	res := 1
	for i := 2; i < n+1; i++ {
		temp = temp * i
		res += temp
	}
	return res
}
