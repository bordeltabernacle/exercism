package letter

type FreqMap map[rune]int

func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan string)
	m := FreqMap{}
	go populateChan(c, texts)
	for s := range c {
		for _, r := range s {
			m[r]++
		}
	}
	return m
}

func populateChan(c chan string, texts []string) {
	for _, t := range texts {
		c <- t
	}
	close(c)
}

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
