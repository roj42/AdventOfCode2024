package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var regexString = `mul\(\d{1,3},\d{1,3}\)`
var regexString2 = `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`

func day3_2(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	//fancy scanner iteration
	do := true
	for scanner.Scan() {
		// log("line:", scanner.Text())
		re := regexp.MustCompile(regexString2)
		for _, match := range re.FindAllString(scanner.Text(), -1) {
			if match == "don't()" {
				do = false
				continue
			}
			if match == "do()" {
				do = true
				continue
			}
			if do {
				grandTotal += warehouseMultiply(match)
			}
		}

	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces if there is one, so yay
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fmt.Sprint(grandTotal)
}

func day3(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	//fancy scanner iteration
	for scanner.Scan() {
		// log("line:", scanner.Text())
		re := regexp.MustCompile(regexString)
		for _, match := range re.FindAllString(scanner.Text(), -1) {
			grandTotal += warehouseMultiply(match)
		}

	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces if there is one, so yay
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fmt.Sprint(grandTotal)
}

func warehouseMultiply(input string) int {
	//mul(xxx,yyy)
	strip := parenSplit(input)
	//xxx,yyy
	values := strings.Split(strip, ",")
	// [xxx,yyy]
	x, err := strconv.Atoi(values[0])
	check(err)
	y, err := strconv.Atoi(values[1])
	check(err)
	return x * y

}

func parenSplit(s string) string {
	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s, ")")
		if j >= 0 {
			return s[i+1 : j]
		}
	}
	return ""
}
