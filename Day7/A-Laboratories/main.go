package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "../example_input.txt"

	buf, err := os.Open(path)
	handleErr(err)
	defer buf.Close()
	scanner := bufio.NewScanner(buf)

	scanner.Scan()
	firstRow := scanner.Text()
	maxAsString := strings.Repeat("1", len(firstRow))
	max, err := strconv.ParseInt(maxAsString, 2, 64)

	prevRow := binary(firstRow)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		currRow := binary(line)
		hitSplitters := currRow & prevRow
		sum += countOnes(hitSplitters)
		splitLasers := (hitSplitters << 1) | (hitSplitters >> 1)

		mergedLasers := (prevRow | splitLasers) ^ hitSplitters

		finalRow := mergedLasers & max

		// fmt.Println("------------------------------------")
		// fmt.Println(padLeft(strconv.FormatInt(prevRow, 2), len(firstRow)), " - Previous Row")
		// fmt.Println(padLeft(strconv.FormatInt(currRow, 2), len(firstRow)), " - Current Row")
		// fmt.Println(padLeft(strconv.FormatInt(hitSplitters, 2), len(firstRow)), " - Hit Splitters")
		// fmt.Println(padLeft(strconv.FormatInt(splitLasers, 2), len(firstRow)), " - Split Lasers")
		fmt.Println(padLeft(strconv.FormatInt(finalRow, 2), len(firstRow)), " - Final Row")

		prevRow = finalRow
	}

	fmt.Println("There were", sum, "splits")
}

func countOnes(val int64) int {
	// There is a O(1) way to count the 1s in a binary number, but I don't understand
	// it and won't use it. https://stackoverflow.com/a/17498333
	ones := 0
	for val != 0 {
		val = val & (val - 1)
		ones++
	}
	return ones
}

func binary(line string) int64 {
	binString := make([]rune, len(line))

	for i, char := range line {
		if char == '.' {
			binString[i] = '0'
		} else {
			binString[i] = '1'
		}
	}

	val, err := strconv.ParseInt(string(binString), 2, 64)
	handleErr(err)

	return val
}

func padLeft(message string, length int) string {
	return strings.Repeat("0", length-len(message)) + message
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
