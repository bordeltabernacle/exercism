package pangram

import "strings"

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const testVersion = 1

func IsPangram(sentence string) bool {
	sentence = strings.ToLower(sentence)
	for _, letter := range alphabet {
		if strings.Contains(sentence, string(letter)) == false {
			return false
		}
	}
	return true
}
