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
	c := make(chan FreqMap)
	m := FreqMap{}

	go func(c chan FreqMap, texts []string) {
		for _, t := range texts {
			c <- Frequency(t)
		}
		close(c)
	}(c, texts)

	for s := range c {
		for letter, count := range s {
			m[letter] += count
		}
	}

	return m
}
