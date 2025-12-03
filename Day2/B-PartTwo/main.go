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
	scanner.Split(scanCommas)

	sum := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		sum += sumInvalidIdsInRange(parts[0], parts[1])
	}
	fmt.Println("Final result is", sum)
}

// For the given range (represented as number strings), finds all invalid ids within the range and returns the sum of those invalid ids
// (I'm just going to brute force this one.)
func sumInvalidIdsInRange(minS string, maxS string) int {
	minI, _ := strconv.Atoi(minS)
	maxI, _ := strconv.Atoi(maxS)

	sum := 0
	for i := minI; i <= maxI; i++ {
		if isInvalid(i) {
			sum += i
		}
	}
	return sum
}

// Not proud of this one. We iterate over every possible id and do all possible even splits. If each part of the split is the same
// then we count it as invalid. Not enough time today to figure out a non-brute force method.
func isInvalid(val int) bool {
	valS := strconv.Itoa(val)

	inputLen := len(valS)
	for parts := 2; parts <= inputLen; parts++ {
		if inputLen%parts != 0 {
			continue
		}

		partLen := inputLen / parts
		firstPart := valS[0:partLen]

		valid := false
		for i := 1; i < parts; i++ {
			start := partLen * i
			end := (i + 1) * partLen
			newPart := valS[start:end]

			if newPart != firstPart {
				valid = true
				break
			}
		}

		if !valid {
			// fmt.Println("Found an invalid id", val)
			return true
		}

	}

	return false
}

// Unnecessary stuff I added to read an input file split by commas instead of newlines

// An implementation of SplitFunc for the scanner which lets us read separated by commas instead of newlines
func scanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
