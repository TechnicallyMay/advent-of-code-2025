package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read in the input file
	path := "../sample_input.txt"
	buf, _ := os.Open(path)
	defer buf.Close()

	scanner := bufio.NewScanner(buf)

	// The dial is supposed to start at 50
	sum := 50
	zeroCount := 0

	// Iterate over every line in the file instead of reading the whole
	// file into memory
	for scanner.Scan() {
		input := scanner.Text()
		direction := string(input[0])
		distance, _ := strconv.Atoi(input[1:])

		if direction == "L" {
			distance *= -1
		}

		sum += distance
		sum = wrap(sum)
		if sum == 0 {
			zeroCount++
		}
	}
	fmt.Println("The dial landed on zero", zeroCount, "times.")
}

// Wrap the given value around the dial, assumes a minimum of 0
// maximum of 100
func wrap(val int) int {
	domain := 100

	if val < 0 {
		val += domain * (-val/100 + 1)
	}

	return val % domain
}
