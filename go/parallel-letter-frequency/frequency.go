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
	m := FreqMap{}
	c := make(chan FreqMap, len(texts))

	for _, t := range texts {
		go func(s string) {
			c <- Frequency(s)
		}(t)
	}

	for i := 0; i < len(texts); i++ {
		freq := <-c
		for k, v := range freq {
			m[k] += v
		}
	}
	return m
}
