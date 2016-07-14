package bob

import (
	"strings"
)

const testVersion = 2

func Hey(s string) (response string) {
	if s == strings.ToUpper(s) {
		return "Whoa, chill out!"
	}
	if strings.HasSuffix(s, "?") {
		return "Sure."
	}
	return "Whatever."
}
