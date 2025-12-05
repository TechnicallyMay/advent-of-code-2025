package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "../input.txt"

	buf, err := os.Open(path)
	handleErr(err)
	defer buf.Close()

	scanner := bufio.NewScanner(buf)

	rows := make([]string, 0)

	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}

	sum := 0
	for true {
		newRows, removed := run(rows)
		sum += removed
		rows = newRows

		if removed == 0 {
			break
		}
	}

	fmt.Println("There were", sum, "available paper rolls")
}

func run(rows []string) (newRows []string, removed int) {
	newRows = make([]string, len(rows))

	defaultRow := strings.Repeat(".", len(rows[0]))
	for i, row := range rows {
		newRow := make([]rune, len(defaultRow))
		prevRow := defaultRow
		nextRow := defaultRow

		if i > 0 {
			prevRow = rows[i-1]
		}

		if i < len(rows)-1 {
			nextRow = rows[i+1]
		}

		for j := range row {
			if isGrabbable(prevRow, row, nextRow, j) {
				removed++
				newRow[j] = '.'
			} else {
				newRow[j] = rune(row[j])
			}
		}
		newRows[i] = string(newRow)
	}

	return
}

var checks [][]int = [][]int{
	{0, 1},  // UP
	{0, -1}, // DOWN

	{-1, 0}, // LEFT
	{1, 0},  // RIGHT

	{1, 1},   // UP-RIGHT
	{1, -1},  // UP-LEFT
	{-1, 1},  // DOWN-RIGHT
	{-1, -1}, // DOWN-LEFT
}

func isGrabbable(prevRow string, currentRow string, nextRow string, charIndex int) bool {
	if currentRow[charIndex] == '.' {
		return false
	}
	rows := []string{
		prevRow,
		currentRow,
		nextRow,
	}

	count := 0
	for _, dir := range checks {
		columnIndexToCheck := charIndex + dir[0]
		rowIndexToCheck := 1 + dir[1]

		if columnIndexToCheck < 0 || columnIndexToCheck >= len(rows[1]) {
			continue
		}

		if rows[rowIndexToCheck][columnIndexToCheck] == '@' {
			count++
		}

		if count == 4 {
			return false
		}
	}

	return true
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
