package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PuzzleRow struct {
	DesiredState int
	Buttons      []int
}

func main() {
	path := "../input.txt"

	buf, err := os.Open(path)
	handleErr(err)
	defer buf.Close()
	scanner := bufio.NewScanner(buf)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += solveLine(line)
	}
	fmt.Println(sum, "button presses were required")
}

func solveLine(line string) int {
	puzzRow := parseLine(line)
	if puzzRow.DesiredState == 0 {
		return 0
	}

	i := 1
	// 2D array of number of button presses => possible numbers formed with that number of presses
	possibilities := make([][]int, 0)
	possibilities = append(possibilities, []int{0})
	for true {
		possibilities = append(possibilities, make([]int, 0))
		prevPossibilities := possibilities[i-1]

		for _, prev := range prevPossibilities {
			for _, butt := range puzzRow.Buttons {
				newPoss := prev ^ butt

				if newPoss == puzzRow.DesiredState {
					return i
					// fmt.Println("To reach the desired state", puzzRow.DesiredState, ",", i, "button presses were required")
					// fmt.Println(possibilities)
					// fmt.Println()
				}

				possibilities[i] = append(possibilities[i], newPoss)
			}
		}

		i++
	}

	panic("This shouldn't have been reached")
}

// Not the prettiest code, but just parses the inputs into binary numbers.
// See ./puzzle.md #Plan
func parseLine(line string) PuzzleRow {
	parts := strings.Split(line, " ")
	lights := removeBraces(parts[0])
	desiredState := 0
	for i, char := range lights {
		if char == '#' {
			// Turn on the light (set the bit to one)
			light := 1 << (len(lights) - i - 1)
			desiredState |= light
		}
	}

	buttons := make([]int, 0)
	for i := 1; i < len(parts); i++ {
		part := parts[i]
		if part[0] == '{' {
			break
		}
		part = removeBraces(part)
		nums := strings.Split(part, ",")

		button := 0
		length := len(lights)
		for _, numStr := range nums {
			lightIndex, err := strconv.Atoi(numStr)
			handleErr(err)
			if lightIndex >= length {
				// A button which toggles a non-existent light
				continue
			}

			lightBit := 1 << (length - lightIndex - 1)
			button |= lightBit

		}
		buttons = append(buttons, button)
	}

	return PuzzleRow{DesiredState: desiredState, Buttons: buttons}
}

func removeBraces(text string) string {
	return text[1 : len(text)-1]
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
