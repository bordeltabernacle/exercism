package house

import "fmt"

var verses = []string{
	"the horse and the hound and the horn\nthat belonged to",
	"the farmer sowing his corn\nthat kept",
	"the rooster that crowed in the morn\nthat woke",
	"the priest all shaven and shorn\nthat married",
	"the man all tattered and torn\nthat kissed",
	"the maiden all forlorn\nthat milked",
	"the cow with the crumpled horn\nthat tossed",
	"the dog\nthat worried",
	"the cat\nthat killed",
	"the rat\nthat ate",
	"the malt\nthat lay in",
}

func Embed(pre, post string) string {
	return fmt.Sprintf("%s %s", pre, post)
}

func Verse(pre string, verses []string, post string) string {
	if len(verses) == 0 {
		return Embed(pre, post)
	}
	return Embed(pre, Verse(verses[0], verses[1:], post))
}

func Song() (song string) {
	first := "This is"
	last := "the house that Jack built."
	for i := 0; i <= len(verses); i++ {
		start := len(verses) - i
		end := len(verses)
		song += Verse(first, verses[start:end], last)
		if i != len(verses) {
			song += "\n\n"
		}
	}
	return
}
