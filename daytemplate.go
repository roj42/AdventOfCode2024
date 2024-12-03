package main

import (
	"bufio"
	"fmt"
)

func dayt(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	//fancy scanner iteration
	for scanner.Scan() {
		// log("line:", scanner.Text())

	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces if there is one, so yay
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fmt.Sprint(grandTotal)
}
