package pangram

import "unicode"

const (
	testVersion = 1
	allChars    = 0x03ffffff
)

func IsPangram(s string) bool {
	var set int32

	for _, r := range s {
		if lr := unicode.ToLower(r); lr >= 'a' && lr <= 'z' {
			set |= 1 << (byte(lr) - 'a')
		}
	}

	return allChars == set
}
