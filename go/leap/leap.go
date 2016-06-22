package leap

const TestVersion = 1

// return true if the given year is a leap year,
// otherwise return false
func IsLeapYear(year int) bool {
	switch {
	case year%4 != 0:
		return false
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	default:
		return true
	}
}
