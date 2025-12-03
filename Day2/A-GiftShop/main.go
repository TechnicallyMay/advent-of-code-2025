package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Answer for my input was 21898734247

func main() {
	path := "../input.txt"
	buf, _ := os.Open(path)
	defer buf.Close()

	scanner := bufio.NewScanner(buf)
	scanner.Split(ScanCommas)

	sum := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		sum += sumInvalidIdsInRange(parts[0], parts[1])
	}
	fmt.Println("Final result is", sum)
}

// For the given range (represented as number strings), finds all invalid ids within the range and returns the sum of those invalid ids
func sumInvalidIdsInRange(minS string, maxS string) int {
	rangeMax, maxerr := strconv.Atoi(maxS)
	rangeMin, minerr := strconv.Atoi(minS)

	if maxerr != nil {
		panic(maxerr)
	} else if minerr != nil {
		panic(minerr)
	}

	if rangeMin > rangeMax {
		panic(fmt.Sprintf("%s - %s or as nums %d - %d", minS, maxS, rangeMin, rangeMax))
	}
	curr := minS
	sum := 0

	for true {
		leftHalfString, leftHalfInt, rightHalfString, rightHalfInt := splitNumberString(curr)

		match := 0
		if len(curr)%2 != 0 {
			// Do nothing, we can't have a match if the string can't be evenly split
		} else if leftHalfInt == rightHalfInt {
			match, _ = strconv.Atoi(leftHalfString + rightHalfString)
		} else if leftHalfInt > rightHalfInt {
			match, _ = strconv.Atoi(leftHalfString + leftHalfString)
		}

		if match > rangeMax {
			break
		}

		sum += match
		curr = strconv.Itoa(leftHalfInt+1) + strings.Repeat("0", len(rightHalfString))
	}

	return sum
}

// Given a string containing a number, splits it in half. Returns the string and integer representation of each half
func splitNumberString(val string) (leftS string, leftI int, rightS string, rightI int) {
	length := len(val) / 2

	leftS = val[:length]
	leftI, _ = strconv.Atoi(leftS)

	rightS = val[length:]
	rightI, _ = strconv.Atoi(rightS)

	return
}

// Unnecessary stuff I added to read an input file split by commas instead of newlines

// An implementation of SplitFunc for the scanner which lets us read separated by commas instead of newlines
func ScanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, dropNewline(data[0:i]), nil

	}

	if atEOF {
		return len(data), dropNewline(data), nil
	}

	// Request more data.
	return 0, nil, nil
}

// drop a \n from the data.
func dropNewline(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\n' {
		return data[0 : len(data)-1]
	}
	return data
}
