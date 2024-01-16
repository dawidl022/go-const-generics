package reversed

func reversedSlice(s []int) []int {
	n := len(s)
	newS := make([]int, n)
	for i := 0; i < n/2; i++ {
		newS[i], newS[n-i-1] = s[n-i-1], s[i]
	}
	return newS
}
