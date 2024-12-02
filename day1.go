package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// compare two lists, multiplying value of each left list item by the number of times it appears in the right list. Sum these.
func day1_2(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	//Our Data. Counts of things? Sounds like a map!
	leftMap := make(map[int]int)
	rightMap := make(map[int]int)

	//fancy scanner iteration
	for scanner.Scan() {
		// log("line:", scanner.Text())

		line := strings.Split(scanner.Text(), "   ")

		//update map for the left item
		leftNum, err := strconv.Atoi(line[0])
		check(err)
		leftMap[leftNum] = leftMap[leftNum] + 1

		rightNum, err := strconv.Atoi(line[1])
		check(err)
		rightMap[rightNum] = rightMap[rightNum] + 1

	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces if there is one, so yay
	if err := scanner.Err(); err != nil {
		check(err)
	}

	//go through the lists, multiplying matches by counts.
	for keyNum := range leftMap {
		//multiply number of times keyNum appears in right by itself, also by how many times lefykey appeared
		// log("key ", keyNum, ": ", " left times ", leftMap[keyNum], ", right times ", rightMap[keyNum])
		grandTotal += (keyNum * leftMap[keyNum] * rightMap[keyNum])
	}

	return fmt.Sprint(grandTotal)
}

// compare two lists, calculating the diff between the two, adding |diff| to a total.
func day1(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	//Our Data
	leftList := []int{}
	rightList := []int{}

	//fancy scanner iteration
	for scanner.Scan() {
		// log("line:", scanner.Text())

		//cram each side of each line into it's slice
		line := strings.Split(scanner.Text(), "   ")
		leftNum, err := strconv.Atoi(line[0])
		check(err)
		leftList = append(leftList, leftNum)
		rightNum, err := strconv.Atoi(line[1])
		check(err)
		rightList = append(rightList, rightNum)

	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces if there is one, so yay
	if err := scanner.Err(); err != nil {
		check(err)
	}

	//sort the lines
	sort.Ints(leftList)
	sort.Ints(rightList)

	//go through the lists, comparing distances and adding to our grand total
	for i := range leftList {
		distance := absDiff(leftList[i], rightList[i])
		// log("Line ", i, ": ", distance, " between ", leftList[i], " and ", rightList[i])
		grandTotal += distance
	}

	return fmt.Sprint(grandTotal)
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
