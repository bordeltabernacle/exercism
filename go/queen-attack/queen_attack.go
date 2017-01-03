package queenattack

import "fmt"

const testVersion = 1

func CanQueenAttack(w, b string) (attack bool, err error) {
	if err = validSquare(w); err != nil {
		return false, err
	}
	if err = validSquare(b); err != nil {
		return false, err
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

func validSquare(s string) error {
	if len(s) != 2 || s[0] > 'h' || s[1] > '7' {
		return fmt.Errorf("Invalid square: %s", s)
	}
	return nil
}
