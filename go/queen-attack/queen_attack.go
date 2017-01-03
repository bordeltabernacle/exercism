package queenattack

import (
	"errors"
	"fmt"
)

const testVersion = 1

// CanQueenAttack ...
func CanQueenAttack(w, b string) (attack bool, err error) {
	if w == b {
		return false, errors.New("Pieces cannot be in same location")
	}
	wrow, wcol, err := getRowAndCol(w)
	if err != nil {
		return false, err
	}
	brow, bcol, err := getRowAndCol(b)
	if err != nil {
		return false, err
	}
	if wrow == brow || wcol == bcol {
		return true, nil
	}
	rowDiff := brow - wrow
	if rowDiff < 0 {
		rowDiff = rowDiff * -1
	}
	colDiff := bcol - wcol
	if colDiff < 0 {
		colDiff = colDiff * -1
	}
	if rowDiff == colDiff {
		return true, nil
	}
	return
}

func getRowAndCol(loc string) (row int, col int, err error) {
	if len(loc) != 2 {
		return 0, 0, fmt.Errorf("%s is an invalid location", loc)
	}
	if loc[0] > 'h' || loc[1] > '7' {
		return 0, 0, fmt.Errorf("%s is an invalid location", loc)
	}
	return int(loc[0]) - 97, int(loc[1]) - 48, nil

}
