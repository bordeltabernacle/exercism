package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2

func Hey(s string) (response string) {
	if strings.TrimSpace(s) == "" {
		return "Fine. Be that way!"
	}
	for _, c := range s {
		if unicode.IsLetter(c) {
			if s == strings.ToUpper(s) {
				return "Whoa, chill out!"
			}
		}
	}
	if strings.HasSuffix(strings.TrimRight(s, " "), "?") {
		return "Sure."
	}
	return "Whatever."
}
