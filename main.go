package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// don't do this in real life
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// don't do this in real life
// actually this is PROBABLY kind of okay.
func log(inputs ...any) {
	fmt.Println(inputs...)
}

func main() {

	//flags for test and second puzzle

	dayPtr := flag.Int("d", 1, "the day")
	testPtr := flag.Bool("t", false, "test file?")
	secondPtr := flag.Bool("2", false, "second puzzle?")

	flag.Parse()

	//day string
	dayInput := strconv.Itoa(*dayPtr)

	filePrefix := dayInput

	//determine if second
	if *secondPtr {
		dayInput = dayInput + "_2"
	}

	//testflag means load the test file
	if *testPtr {
		filePrefix = filePrefix + "_t"
	}
	start := time.Now()

	//open the day's file, and close it when we're done with main, here.
	fileName := "./data/" + filePrefix + ".txt"

	file, err := os.Open(fileName)
	check(err)
	log("opened " + fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, you can resize scanner's capacity for lines over 64K? I hope I never need this note.

	//WHAT DAY IS IT?
	switch dayInput {
	case "1":
		fmt.Println(day1(scanner))
	case "1_2":
		fmt.Println(day1_2(scanner))
	case "2":
		fmt.Println(day2(scanner))
	case "2_2":
		fmt.Println(day2_2(scanner))
	case "3":
		fmt.Println(day3(scanner))
	case "3_2":
		fmt.Println(day3_2(scanner))
	case "4":
		fmt.Println(day4(scanner))

	default:
		log("no implementation for day: " + dayInput)

	}
	stop := time.Since(start)
	log("time", stop.String())

}
