package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read in the input file
	path := "../input.txt"
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

			// This is a silly correction I have to make. My algorithm doesn't count when we turn left and land on zero, instead it counts
			// when we wrap around to 99. So, preempt this by subtracting the wrap it's going to count.
			// I'd like to come up with a more elegant solution.
			if sum == 0 {
				zeroCount--
			}
		}

		sum += distance
		zeroPasses := 0
		sum, zeroPasses = wrap(sum)
		zeroCount += zeroPasses

		// Part of the silly correction as well. It doesn't count when we land on zero since we didn't technically wrap.
		if direction == "L" && sum == 0 {
			zeroCount += 1
		}
	}
	fmt.Println("The dial landed on zero", zeroCount, "times.")
}

// Wrap the given value around the dial, assumes a minimum of 0
// maximum of 100. Count how many times we wrap
func wrap(val int) (int, int) {
	orig := val
	domain := 100

	if val < 0 {
		val += domain * (-val/domain + 1)
	}

	result := val % domain
	zeros := abs((result - orig) / domain)

	return result, zeros
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}
