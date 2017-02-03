package queenattack

import "fmt"

const testVersion = 1

func CanQueenAttack(w, b string) (attack bool, err error) {
	if !validSquare(w) || !validSquare(b) {
		return false, fmt.Errorf("Invalid square: %s", w+b)
	}
	if w == b {
		return false, fmt.Errorf("Queens cannot be on same square: %s", w)
	}
	if w[0] == b[0] || w[1] == b[1] {
		return true, nil
	}
	fileDiff := w[0] - b[0]
	rankDiff := w[1] - b[1]
	return fileDiff == rankDiff || fileDiff+rankDiff == 0, nil
}

func validSquare(s string) bool {
	return !(len(s) != 2 ||
		s[0] > 'h' || s[0] < 'a' ||
		s[1] > '7' || s[1] < '1')
}
