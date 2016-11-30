package foodchain

import "fmt"

var testVersion = 2

var wriggle = "wriggled and jiggled and tickled inside her"

var eaten = []struct {
	beast   string
	comment string
}{
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It " + wriggle + "."},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
	{"horse", "She's dead, of course!"},
}

func Verse(v int) (verse string) {
	v--
	verse = fmt.Sprintf("I know an old lady who swallowed a %s.\n%s", eaten[v].beast, eaten[v].comment)
	if v > 0 && v < 7 {
		for i := v; i > 0; i-- {
			verse += fmt.Sprintf("\nShe swallowed the %s to catch the %s", eaten[i].beast, eaten[i-1].beast)
			if i == 2 {
				verse += fmt.Sprintf(" that %s", wriggle)
			}
			verse += fmt.Sprintf(".")
		}
		verse += fmt.Sprintf("\n%s", eaten[0].comment)
	}
	return
}

func Verses(start, end int) (verses string) {
	for i := start; i <= end; i++ {
		if i > 1 {
			verses += "\n\n"
		}
		verses += fmt.Sprintf("%s", Verse(i))
	}
	return
}

func Song() (song string) {
	return Verses(1, 8)
}
