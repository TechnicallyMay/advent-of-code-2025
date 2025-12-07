package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	path := "../input.txt"

	buf, err := os.Open(path)
	handleErr(err)
	defer buf.Close()
	scanner := bufio.NewScanner(buf)

	operands := make([]string, 1)
	operators := make([]rune, 1)
	operatorIndexes := make([]int, 1)

	for scanner.Scan() {
		line := scanner.Text()

		for charIndex, char := range line {
			if char == '+' || char == '*' {
				operators = append(operators, char)
				operatorIndexes = append(operatorIndexes, charIndex)
				continue
			}

			if charIndex >= len(operands) {
				operands = append(operands, "")
			}

			if char == ' ' {
				continue
			}

			operands[charIndex] += string(char)
		}
	}

	results := make([]int, len(operators))

	operatorIndex := 0
	for i, stringOp := range operands {
		op, err := strconv.Atoi(stringOp)
		if err != nil {
			continue
		}

		if operatorIndex < len(operators)-1 && i >= operatorIndexes[operatorIndex+1] {
			operatorIndex++

			results[operatorIndex] = op
			continue
		}

		switch operators[operatorIndex] {
		case '+':
			results[operatorIndex] += op
		case '*':
			results[operatorIndex] *= op
		}

	}

	sum := 0
	for _, result := range results {
		sum += result
	}
	fmt.Println(sum)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
