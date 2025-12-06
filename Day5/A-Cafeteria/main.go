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

	buf, err := os.Open(path)
	handleErr(err)
	defer buf.Close()
	scanner := bufio.NewScanner(buf)

	// Pairs of ints representing min/max in a range
	ranges := make([][]int, 0)
	// Parse the lines of the file to find the ranges
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			break
		}

		vals := strings.Split(row, "-")
		min, err := strconv.Atoi(vals[0])
		handleErr(err)
		max, err := strconv.Atoi(vals[1])
		handleErr(err)

		ranges = append(ranges, []int{min, max})
	}

	sum := 0
	// Now that we have the ranges, keep scanning for the ids of fresh veg
	for scanner.Scan() {
		row := scanner.Text()
		val, err := strconv.Atoi(row)
		handleErr(err)
		sum += countFresh(val, ranges)
	}

	fmt.Println("There were", sum, "fresh")
}

func countFresh(val int, ranges [][]int) int {
	for _, r := range ranges {
		if val >= r[0] && val <= r[1] {
			return 1
		}
	}
	return 0
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
