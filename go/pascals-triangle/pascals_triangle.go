package pascal

func Triangle(n int) (triangle [][]int) {
	for i := 0; i < n; i++ {
		triangle = append(triangle, line(i))
	}
	return
}

func line(n int) []int {
	l := []int{1}
	for i := 0; i < n; i++ {
		l = append(l, l[i]*(n-i)/(i+1))
	}
	return l
}
