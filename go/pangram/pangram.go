package pangram

import "strings"

const testVersion = 1

func IsPangram(sentence string) bool {
	sentence = strings.ToLower(sentence)
	for i := 97; i < 123; i++ {
		if strings.Contains(sentence, string(i)) == false {
			return false
		}
	}
	return true
}
