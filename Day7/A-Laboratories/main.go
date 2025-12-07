package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	path := "../input.txt"

	buf, err := os.Open(path)
	handleErr(err)
	defer buf.Close()
	scanner := bufio.NewScanner(buf)

	scanner.Scan()
	firstRow := scanner.Text()
	maxAsString := strings.Repeat("1", len(firstRow))

	max := parseBin(maxAsString)

	prevRow := binary(firstRow)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		currRow := binary(line)

		hitSplitters := new(big.Int)
		hitSplitters.And(currRow, prevRow)
		sum += countOnes(hitSplitters)

		splitLasers := new(big.Int)
		leftShift := new(big.Int)
		rightShift := new(big.Int)
		leftShift.Lsh(hitSplitters, 1)
		rightShift.Rsh(hitSplitters, 1)
		splitLasers.Or(leftShift, rightShift)

		mergedLasers := new(big.Int)
		mergedLasers.Or(prevRow, splitLasers)

		mergedLasersWithoutSplitters := new(big.Int)
		mergedLasersWithoutSplitters.Xor(mergedLasers, hitSplitters)

		finalRow := new(big.Int)
		finalRow.And(mergedLasersWithoutSplitters, max)

		// fmt.Println("------------------------------------")
		// fmt.Printf("%b - Prev\n", prevRow)
		// fmt.Printf("%b - Curr\n", currRow)
		// fmt.Printf("%b - Hit\n", hitSplitters)
		// fmt.Printf("%b - Split\n", splitLasers)
		// fmt.Printf("%b - Merged\n", mergedLasers)
		// fmt.Printf("%b - Merged - Split\n", mergedLasersWithoutSplitters)
		fmt.Printf("%b - Final Row\n", finalRow)
		prevRow = finalRow
	}

	fmt.Println("There were", sum, "splits")
}

func countOnes(val *big.Int) int {
	// There is a O(1) way to count the 1s in a binary number, but I don't understand
	// it and won't use it. https://stackoverflow.com/a/17498333
	ones := 0

	copy := new(big.Int)
	copy.Set(val)
	zero := new(big.Int)
	one := new(big.Int)
	one.SetBit(one, 0, 1)
	for copy.Cmp(zero) != 0 {
		copyM1 := new(big.Int)
		copyM1.Sub(copy, one)

		copy.And(copy, copyM1)
		ones++
	}
	return ones
}

func binary(line string) *big.Int {
	binString := make([]rune, len(line))

	for i, char := range line {
		if char == '.' {
			binString[i] = '0'
		} else {
			binString[i] = '1'
		}
	}

	return parseBin(string(binString))
}

func parseBin(bin string) *big.Int {
	i := new(big.Int)
	_, err := fmt.Sscanf(string(bin), "%b", i)
	handleErr(err)

	return i
}

func padLeft(message string, length int) string {
	return strings.Repeat("0", length-len(message)) + message
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
