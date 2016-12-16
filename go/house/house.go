package house

import "fmt" s


var lines = []string{
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
	"the house that Jack built.",
}

func Embed(relPhrase, nounPhrase string) string {
	return fmt.Sprintf("%s %s", relPhrase, nounPhrase)
}

func Verse(subject string, relPhrases []string, nounPhrase string) string {
	if len(relPhrases) == 0 {
		return Embed(subject, nounPhrase)
	}
	return Embed(subject, Verse(relPhrases[0], relPhrases[1:], nounPhrase))
}

func Song() (song string) {
	first := "This is"
	last := lines[len(lines)-1]
	for i := 1; i <= len(lines); i++ {
		start := len(lines) - i
		end := len(lines) - 1
		song += Verse(first, lines[start:end], last)
		if i != len(lines) {
			song += "\n\n"
		}
	}
	return
}
