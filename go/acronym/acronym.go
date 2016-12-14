package acronym

import (
	"regexp"
	"strings"
)

const testVersion int = 1

func abbreviate(name string) (abbreviation string) {
	pattern := "[A-Z]+[a-z]*|[a-z]+"
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(name, -1)
	for _, match := range matches {
		firstLetter := strings.Split(match, "")[0]
		abbreviation += strings.ToUpper(firstLetter)
	}
	return
}
