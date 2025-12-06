package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	sortRanges(ranges)
	ensureNonOverlapping(ranges)
	// fmt.Println(ranges)

	sum := 0
	for _, r := range ranges {
		// Have to add one to make the count inclusive
		sum += r[1] - r[0] + 1
	}

	fmt.Println("There were", sum, "fresh")
}

// Sorts the ranges in order of their min value
func sortRanges(ranges [][]int) {
	slices.SortFunc(ranges, func(a, b []int) int {
		return a[0] - b[0]
	})
}

func ensureNonOverlapping(ranges [][]int) {
	for i := 0; i < len(ranges)-1; i++ {
		currRange := ranges[i]
		nextRange := ranges[i+1]

		if nextRange[1] < currRange[1] {
			nextRange[1] = currRange[1]
		}

		if currRange[1] >= nextRange[0] {
			currRange[1] = nextRange[0] - 1
		}
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
