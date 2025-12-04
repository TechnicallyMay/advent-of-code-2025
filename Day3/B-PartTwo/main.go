package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Answer for example_input is 3121910778619
func main() {
	path := "../input.txt"
	buf, err := os.Open(path)
	handleErr(err)

	defer buf.Close()

	scanner := bufio.NewScanner(buf)

	var sum int = 0
	for scanner.Scan() {
		input := scanner.Text()
		res := solve(input)
		sum += res
		fmt.Println("Result was", res, "with input", input)
	}

	fmt.Println("Final result is", sum)
}

// Finds the largest possible number that can be formed using 12 digits from the string.
// Digits must go from left to right
func solve(input string) int {
	digits := make([]rune, 0)

	end := len(input)
	startIndex := 0
	for i := range 12 {
		remainingDigits := 12 - i
		endIndex := end - remainingDigits + 1
		val, index := findLargestDigitInRange(input, startIndex, endIndex)
		startIndex = index + 1

		digits = append(digits, val)
	}

	res, err := strconv.Atoi(string(digits))
	handleErr(err)

	return res
}

func findLargestDigitInRange(input string, start int, end int) (val rune, index int) {
	runes := []rune(input)
	var max int32 = 0
	for i := start; i < end; i++ {
		// Magic - subtracting the char '0' from any rune 0-9 gives the int rep of that rune
		int := int32(input[i] - '0')
		if int > max {
			max = int
			val = runes[i]
			index = i
		}
		if int == 9 {
			break
		}
	}

	return
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
