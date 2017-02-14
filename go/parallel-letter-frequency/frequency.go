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

	for _, t := range texts {
		go func(t string) {
			c <- Frequency(t)
		}(t)
	}

	for range texts {
		for k, v := range <-c {
			m[k] += v
		}
	}
	return m
}
