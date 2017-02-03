package pangram

import "unicode"

const testVersion = 1

func IsPangram(s string) bool {
	if len(sentence) < 26 {
		return false
	}
	alphabet := map[rune]int{}
	for _, r := range sentence {
		lr := unicode.ToLower(r)
		if lr >= 'a' && lr <= 'z' {
			alphabet[lr] += 1
		}
	}
	return len(alphabet) == 26
}
