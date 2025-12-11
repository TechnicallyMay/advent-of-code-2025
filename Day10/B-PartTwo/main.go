package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/mxschmitt/golang-combinations"
)

func main() {
	path := "../input.txt"

	buf, err := os.Open(path)
	handleErr(err)
	defer buf.Close()
	scanner := bufio.NewScanner(buf)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		desiredOutput, buttons := parseLine(line)
		strtCombos := generateStartingCombos(desiredOutput, buttons)
		newPresses := solveRecursive(strtCombos, desiredOutput)

		sum += newPresses
	}
	fmt.Println(sum, "button presses were required")
}

func solveRecursive(combos map[int][][]int, desiredOutput []int) int {
	if allZeros(desiredOutput) {
		return 0
	}

	res := 10_000
	for cost := range combos {
		combosForCost := combos[cost]
		for _, combo := range combosForCost {
			if !canPressButton(desiredOutput, combo) {
				continue
			}

			// Press the button(s)
			applied := subtractArrays(desiredOutput, combo)
			newGoal := divideArray(applied, 2)

			newRes := cost + (2 * solveRecursive(combos, newGoal))

			if newRes <= res {
				res = newRes
			}
		}
	}
	return res
}

func canPressButton(desiredOutput []int, button []int) bool {
	for i := range desiredOutput {
		b := button[i]
		o := desiredOutput[i]
		if b <= o && b%2 == o%2 {
			continue
		}
		return false
	}

	return true
}

// Produces every (unique) combination which can be formed with the given buttons. Pressing a button twice
// Effectively does nothing
func generateStartingCombos(desiredOutput []int, buttons [][]int) map[int][][]int {
	res := make(map[int][][]int, 0)

	// We need to be able to consider "not pressing a button" as a valid solution to the top level problem
	// i.e. if my desired output is 10 and my buttons are 1 OR 5 the optimal result is to press nothing, then split
	// the problem in half and press 5 twice. If we don't include 0, then it would have to press 1 firt
	res[0] = make([][]int, 1)
	res[0][0] = make([]int, len(desiredOutput))

	all := make([][]int, 0)
	for i := 1; i <= len(buttons); i++ {
		combos := combinations.Combinations(buttons, i)

		for _, combo := range combos {
			sum := make([]int, len(desiredOutput))
			for _, button := range combo {
				sum = addArrays(sum, button)
			}
			dupe := false
			for i := range all {
				if arraysEqual(all[i], sum) {
					dupe = true
					break
				}
			}

			if dupe {
				continue
			}
			all = append(all, sum)
			res[i] = append(res[i], sum)
		}
	}

	return res
}

func arraysEqual(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func allZeros(a []int) bool {
	for _, val := range a {
		if val != 0 {
			return false
		}
	}
	return true
}

func addArrays(a, b []int) []int {
	res := make([]int, len(a))
	for i := range a {
		res[i] = a[i] + b[i]
	}
	return res
}

func subtractArrays(a, b []int) []int {
	res := make([]int, len(a))
	for i := range a {
		res[i] = a[i] - b[i]
	}
	return res
}

func divideArray(a []int, b int) []int {
	res := make([]int, len(a))

	for i := range a {
		res[i] = a[i] / b
	}
	return res
}

// See ./puzzle.md #Plan
func parseLine(line string) (desiredOutput []int, buttons [][]int) {
	// Cut off the first one, which is the light configuration we don't care about anymore
	parts := strings.Split(line, " ")[1:]

	buttons = make([][]int, 0)
	// Stupid workaround to sort tthe 2d array
	is := make([]int, 0)
	desiredOutput = parseNumArray(parts[len(parts)-1], math.MaxInt)

	var i int
	for i = 0; i < len(parts)-1; i++ {
		part := parts[i]
		if part[0] == '{' {
			break
		}
		buttons = append(buttons, numArrayToButton(len(desiredOutput), parseNumArray(part, len(desiredOutput)-1)))
		is = append(is, i)
	}

	return
}

func parseNumArray(text string, max int) []int {
	text = removeBraces(text)
	nums := strings.Split(text, ",")
	res := make([]int, 0)
	for _, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		if num > max {
			continue
		}

		res = append(res, num)
		handleErr(err)
	}

	return res
}

func numArrayToButton(totalLength int, numArray []int) []int {
	button := make([]int, totalLength)
	for _, val := range numArray {
		button[val] = 1
	}
	return button
}

func removeBraces(text string) string {
	return text[1 : len(text)-1]
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
