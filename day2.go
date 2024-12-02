package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// check each line for safety* and add up the number of safe lines
// * a line is safe iff:
// it as all descending or all ascending
func day2(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	//fancy scanner iteration
	for scanner.Scan() {
		// log("line:", scanner.Text())

		line := strings.Split(scanner.Text(), " ")

		//make line into ints
		lineInts := []int{}
		for _, item := range line {
			num, err := strconv.Atoi(item)
			check(err)
			lineInts = append(lineInts, num)
		}

		if safeLine(lineInts) {
			log("safe line: ", scanner.Text())
			grandTotal++
		}
	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces if there is one, so yay
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fmt.Sprint(grandTotal)
}

func day2_2(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	//fancy scanner iteration
out:
	for scanner.Scan() {
		// log("line:", scanner.Text())

		line := strings.Split(scanner.Text(), " ")

		//make line into ints
		lineInts := []int{}
		for _, item := range line {
			num, err := strconv.Atoi(item)
			check(err)
			lineInts = append(lineInts, num)
		}

		if safeLine(lineInts) {
			grandTotal++
			continue out
		}
		//THE PROBLEM DAMPENER - brute force removing one to see if we get a safe.
		for i := 0; i < len(lineInts); i++ {
			//copy to save our old line
			newLine := make([]int, len(lineInts))
			copy(newLine, lineInts)
			newLine = append(newLine[:i], newLine[i+1:]...)
			if safeLine(newLine) {
				grandTotal++
				// log("Safe:", scanner.Text(), "when", newLine)
				continue out
			}
		}
		// log("unsafe", scanner.Text())
	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces if there is one, so yay
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fmt.Sprint(grandTotal)
}

func safeLine(lineInts []int) bool {
	//check line for safety
	var direction bool
	var dirSet bool
	for i := 0; i < len(lineInts)-1; i++ {
		first, second := lineInts[i], lineInts[i+1]

		//unsafe level change or no level change
		if first == second || absDiff(first, second) > 3 {
			return false
		}

		//we can check a direction change
		curDir := first > second
		if !dirSet { //unset direction, establish and skip check
			direction = curDir
			dirSet = true
			continue
		}
		if curDir != direction {
			//direction change!
			return false
		}
	}
	return true
}
