package slice

func All(n int, s string) (substrings []string) {
	if n > len(s) {
		return
	}
	for i := 0; i < len(s)-(n-1); i++ {
		substrings = append(substrings, s[i:i+n])
	}
	return
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return
	}
	return s[:n], true
}
