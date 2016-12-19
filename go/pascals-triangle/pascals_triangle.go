package pascal

func Triangle(n int) (triangle [][]int) {
	for i := 0; i < n; i++ {
		l := []int{1}
		for j := 0; j < i; j++ {
			l = append(l, l[j]*(i-j)/(j+1))
		}
		triangle = append(triangle, l)
	}
	return
}
