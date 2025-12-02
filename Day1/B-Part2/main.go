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

	// Iterate over every line in the file instead of reading the whole file into memory
	for scanner.Scan() {
		input := scanner.Text()
		direction := string(input[0])
		distance, _ := strconv.Atoi(input[1:])

		if direction == "L" {
			distance *= -1
		}

		startingSum := sum
		newSum := startingSum + distance
		sum = wrap(newSum)
		newZeros := calculateZeroPasses(startingSum, newSum, sum)
		zeroCount += newZeros

		// fmt.Println("Starting:", startingSum, "New:", newSum, "Wrapped:", sum, "Zeros:", newZeros)
	}

	fmt.Println("The dial landed on zero", zeroCount, "times.")
}

// Wrap the given value around the dial, assumes a minimum of 0
// maximum of 100. Count how many times we wrap
func wrap(val int) int {
	domain := 100

	if val < 0 {
		val += domain * (-val/domain + 1)
	}

	return val % domain
}

// startingSum - The dial value before turning
// newSum      - The dial value after turning without constraints (can be above 99 and below 0)
// wrappedSum  - The dial value after wrapping between 0-99
func calculateZeroPasses(startingSum int, newSum int, wrappedSum int) int {
	zeros := abs((wrappedSum - newSum) / 100)
	// When we're rotating right (positive number), the number of times we pass zero is
	// equivalent to the number of times we've wrapped (since we wrap from 99 => 0)
	if newSum > 0 {
		return zeros
	}

	// Now make some adjustments for when we're turning left.

	// If we started on 0 we need to subtract a wrap, because the counted wrap
	// didn't actually land us on zero (it was from 0 => 99)
	if startingSum == 0 {
		zeros--
	}

	// If we ended on 0 we need to add a wrap, because it landing on zero doesn't result in a wrap
	if wrappedSum == 0 {
		zeros++
	}

	return zeros

}

// Absolute value for an int
func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}
