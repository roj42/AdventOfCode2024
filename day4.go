package main

import (
	"bufio"
	"fmt"
)

// byteBoard is a two dimensional array of bytes
type byteBoard [][]byte

// isByteAt tells you if the Byte at a certain position is a valid position and is that byte. Empty byte if not
func (b byteBoard) isByteAt(y, x int) byte {
	//nest some ifs for clarity
	//is y valid?
	if y < 0 || y > len(b)-1 {
		return 0
	}
	//is x valid?
	if x < 0 || x > len(b[y])-1 {
		return 0
	}
	return b[y][x]
}

func day4(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	//A 2d slice
	var puzzleBoard byteBoard

	//fancy scanner iteration
	for scanner.Scan() {
		// log("line:", scanner.Text())
		puzzleBoard = append(puzzleBoard, []byte(scanner.Text()))

	}

	//method: scan through for the char X. For each X, look in each direction for M then A then S, stopping if we hit a wall or the wrong letter
	for y, row := range puzzleBoard {
		for x, col := range row {
			if col == 'X' {
				grandTotal += CountXmasFromXAt(&puzzleBoard, y, x)
			}
		}

	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces if there is one, so yay
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fmt.Sprint(grandTotal)
}

func CountXmasFromXAt(puzzleBoard *byteBoard, y, x int) int {
	total := 0
	//Brute it! Start from current point
	//up-left
	if puzzleBoard.isByteAt(y-1, x-1) == 'M' && puzzleBoard.isByteAt(y-2, x-2) == 'A' && puzzleBoard.isByteAt(y-3, x-3) == 'S' {
		total += 1
	}
	//up
	if puzzleBoard.isByteAt(y-1, x) == 'M' && puzzleBoard.isByteAt(y-2, x) == 'A' && puzzleBoard.isByteAt(y-3, x) == 'S' {
		total += 1
	}
	//up-right
	if puzzleBoard.isByteAt(y-1, x+1) == 'M' && puzzleBoard.isByteAt(y-2, x+2) == 'A' && puzzleBoard.isByteAt(y-3, x+3) == 'S' {
		total += 1
	}
	//left
	if puzzleBoard.isByteAt(y, x-1) == 'M' && puzzleBoard.isByteAt(y, x-2) == 'A' && puzzleBoard.isByteAt(y, x-3) == 'S' {
		total += 1
	}
	//right
	if puzzleBoard.isByteAt(y, x+1) == 'M' && puzzleBoard.isByteAt(y, x+2) == 'A' && puzzleBoard.isByteAt(y, x+3) == 'S' {
		total += 1
	}
	//down-left
	if puzzleBoard.isByteAt(y+1, x-1) == 'M' && puzzleBoard.isByteAt(y+2, x-2) == 'A' && puzzleBoard.isByteAt(y+3, x-3) == 'S' {
		total += 1
	}
	//down
	if puzzleBoard.isByteAt(y+1, x) == 'M' && puzzleBoard.isByteAt(y+2, x) == 'A' && puzzleBoard.isByteAt(y+3, x) == 'S' {
		total += 1
	}
	//down-right
	if puzzleBoard.isByteAt(y+1, x+1) == 'M' && puzzleBoard.isByteAt(y+2, x+2) == 'A' && puzzleBoard.isByteAt(y+3, x+3) == 'S' {
		total += 1
	}

	return total

}

func day4_2(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	//A 2d slice
	var puzzleBoard byteBoard

	//fancy scanner iteration
	for scanner.Scan() {
		// log("line:", scanner.Text())
		puzzleBoard = append(puzzleBoard, []byte(scanner.Text()))

	}

	//method: scan through for the char X. For each X, look in each direction for M then A then S, stopping if we hit a wall or the wrong letter
	for y, row := range puzzleBoard {
		for x, col := range row {
			if col == 'A' {
				grandTotal += CountX_MasFromAAt(&puzzleBoard, y, x)
			}
		}

	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces if there is one, so yay
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fmt.Sprint(grandTotal)
}

func CountX_MasFromAAt(puzzleBoard *byteBoard, y, x int) int {
	total := 0
	//Brute it!
	//clock wise! start from...
	//up-left
	if puzzleBoard.isByteAt(y-1, x-1) == 'M' && puzzleBoard.isByteAt(y+1, x+1) == 'S' &&
		puzzleBoard.isByteAt(y-1, x+1) == 'M' && puzzleBoard.isByteAt(y+1, x-1) == 'S' {
		total += 1
	}
	//up-right
	if puzzleBoard.isByteAt(y-1, x+1) == 'M' && puzzleBoard.isByteAt(y+1, x-1) == 'S' &&
		puzzleBoard.isByteAt(y+1, x+1) == 'M' && puzzleBoard.isByteAt(y-1, x-1) == 'S' {
		total += 1
	}
	//down-right
	if puzzleBoard.isByteAt(y+1, x+1) == 'M' && puzzleBoard.isByteAt(y-1, x-1) == 'S' &&
		puzzleBoard.isByteAt(y+1, x-1) == 'M' && puzzleBoard.isByteAt(y-1, x+1) == 'S' {
		total += 1
	}
	//down-left
	if puzzleBoard.isByteAt(y+1, x-1) == 'M' && puzzleBoard.isByteAt(y-1, x+1) == 'S' &&
		puzzleBoard.isByteAt(y-1, x-1) == 'M' && puzzleBoard.isByteAt(y+1, x+1) == 'S' {
		total += 1
	}

	return total

}
