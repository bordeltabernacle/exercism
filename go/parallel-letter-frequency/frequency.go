package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan string)
	m := FreqMap{}

	go func(texts []string, c chan string) {
		defer close(c)
		for _, t := range texts {
			c <- t
		}
	}(texts, c)

	for s := range c {
		for _, r := range s {
			m[r]++
		}
	}
	return m
}
