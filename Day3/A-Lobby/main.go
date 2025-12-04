package main

import (
	"bufio"
	"fmt"
	"os"
)

// Answer for example_input is 357
func main() {
	path := "../input.txt"
	buf, err := os.Open(path)
	handleErr(err)

	defer buf.Close()

	scanner := bufio.NewScanner(buf)

	var sum int32 = 0
	for scanner.Scan() {
		input := scanner.Text()
		res := solve(input)
		sum += res
		// fmt.Println("Result was", res, "with input", input)
	}

	fmt.Println("Final result is", sum)
}

// Finds the largest possible number that can be formed using 2 digits from the string.
// 2nd digit must be to the right of the first digit.
func solve(input string) int32 {

	end := len(input)

	// Only iterate to the 2nd last character of the string for the first one, since the 2nd digit has to be after the first
	first, index := findLargestDigitInRange(input, 0, end-1)
	second, _ := findLargestDigitInRange(input, index+1, end)

	return (first * 10) + second
}

func findLargestDigitInRange(input string, start int, end int) (val int32, index int) {
	for i := start; i < end; i++ {
		// Magic - subtracting the char '0' from any rune 0-9 gives the int rep of that rune
		int := int32(input[i] - '0')
		if int > val {
			val = int
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
