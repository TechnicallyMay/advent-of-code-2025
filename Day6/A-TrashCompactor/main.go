package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "../input.txt"

	lastLine := getLastLineOfFile(path)
	operators := getOperators(lastLine)

	buf, err := os.Open(path)
	handleErr(err)
	defer buf.Close()
	scanner := bufio.NewScanner(buf)

	scanner.Scan()
	// Start with the first line of numbers
	results := getNumbers(scanner.Text())

	for scanner.Scan() {
		line := scanner.Text()
		operands := getNumbers(line)

		for i, operand := range operands {
			operator := operators[i]

			if operator == '+' {
				results[i] += operand
			} else if operator == '*' {
				results[i] *= operand
			}
		}
	}

	sum := 0

	for _, res := range results {
		sum += res
	}

	// fmt.Println(results)
	fmt.Println(sum)
}

func getLastLineOfFile(fname string) string {
	buf, err := os.Open(fname)
	handleErr(err)
	defer buf.Close()

	// Assume the last line is longer than 2 bytes to make scanning easier
	var cursor int64 = -2
	bytes := make([]byte, 1)
	for true {
		_, err := buf.Seek(cursor, 2)
		handleErr(err)
		_, err = buf.Read(bytes)
		handleErr(err)

		if string(bytes) == "\n" {
			break
		}
		cursor--
	}

	lastLineBytes := make([]byte, cursor*-1)
	_, err = buf.Read(lastLineBytes)
	handleErr(err)

	return strings.TrimSpace(string(lastLineBytes))
}

func getOperators(text string) []rune {
	result := make([]rune, 0)

	for _, char := range text {
		if char == '*' || char == '+' {
			result = append(result, char)
		}
	}
	return result
}

func getNumbers(text string) []int {
	result := make([]int, 0)

	parts := strings.Split(text, " ")

	for _, part := range parts {
		int, err := strconv.Atoi(part)
		if err == nil {
			result = append(result, int)
		}
	}

	return result
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
