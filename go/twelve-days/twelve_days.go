package twelve

import "fmt"

const testVersion = 1

var days = []struct {
	day     string
	present string
}{
	{"first", "a Partridge in a Pear Tree."},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

func Verse(v int) (verse string) {
	v--
	verse = fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s", days[v].day, days[v].present)

	if v == 0 {
		return
	}

	for ; v > 0; v-- {
		if v == 1 {
			verse += fmt.Sprintf(", and %s", days[v-1].present)
		} else {
			verse += fmt.Sprintf(", %s", days[v-1].present)
		}
	}
	return
}

func Song() (song string) {
	for i := 1; i <= 12; i++ {
		song += fmt.Sprintf("%s\n", Verse(i))
	}
	return
}
