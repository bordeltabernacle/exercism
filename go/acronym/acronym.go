package acronym

import (
	"regexp"
	"strings"
)

const testVersion int = 1

func abbreviate(name string) (abbreviation string) {
	pattern := "[A-Z]+[a-z]*|[a-z]+"
	re := regexp.MustCompile(pattern)
	match := re.FindAllString(name, -1)
	for _, m := range match {
		abbreviation += strings.ToUpper(m[:1])
	}
	return
}
