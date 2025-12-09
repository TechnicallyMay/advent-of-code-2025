package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vec2 struct {
	x, y int
}

func main() {
	path := "../input.txt"

	buf, err := os.Open(path)
	handleErr(err)
	defer buf.Close()
	scanner := bufio.NewScanner(buf)

	inputs := make([]Vec2, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		handleErr(err)
		y, err := strconv.Atoi(parts[1])
		handleErr(err)

		inputs = append(inputs, Vec2{x: x, y: y})
	}

	maxArea := 0
	for i := 0; i < len(inputs)-1; i++ {
		a := inputs[i]
		for j := i + 1; j < len(inputs); j++ {
			b := inputs[j]
			area := Area(a, b)

			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Println(maxArea)
}

func Area(a Vec2, b Vec2) int {
	height := Abs(a.y-b.y) + 1
	width := Abs(a.x-b.x) + 1

	return height * width
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
