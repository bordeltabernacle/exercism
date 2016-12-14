package pangram

import "strings"

const testVersion = 1

func IsPangram(sentence string) bool {
	sentence = strings.ToLower(sentence)
	for i := int('a'); i <= int('z'); i++ {
		if !strings.Contains(sentence, string(i)) {
			return false
		}
	}
	return true
}
