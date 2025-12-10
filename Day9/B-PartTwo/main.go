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

func (v *Vec2) String() string {
	// Format such that it can be pasted into Desmos as a table
	return fmt.Sprintf("%v\t%v", v.x, v.y)
}

// Correct answer is 1571016172 for my input
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
	// Attempt to form every rect
	for i := range inputs {
		a := inputs[i]
		for j := i + 1; j < len(inputs); j++ {
			b := inputs[j%len(inputs)]
			area := Area(a, b)

			// No need to calculate if the rect is valid until we know it's bigger than the current biggest
			if area > maxArea && IsValidRect(a, b, inputs) {
				maxArea = area
			}
		}
	}

	fmt.Println("Max area was", maxArea)
}

func Area(a Vec2, b Vec2) int {
	height := Abs(a.y-b.y) + 1
	width := Abs(a.x-b.x) + 1

	return height * width
}

// Determines if the formed rectangle intersects any of the lines
// This technically isn't fully correct - it could allow a rect which is formed outside
// of the region as opposed to inside. Luckily, the input is a circle.
func IsValidRect(a, b Vec2, polygon []Vec2) bool {
	paths := MakeRect(a, b)
	for _, path := range paths {
		if Raycast(path[0], path[1], polygon) {
			// msg := "Shape was invalid\n"
			// PrintShape(paths, i)
			// fmt.Println(msg)
			return false
		}
	}

	return true
}

// Given 2 opposite corners of the rectangle, returns all lines that need to be checked for intersections.
// This includes 1 line for each side of the rect, and 2 diagonal lines across the corners
func MakeRect(a, b Vec2) [][]Vec2 {
	oppositeA, oppositeB := Vec2{x: b.x, y: a.y}, Vec2{x: a.x, y: b.y}
	return [][]Vec2{
		{a, b},
		{a, oppositeB},
		{a, oppositeA},
		{oppositeA, oppositeB},
		{oppositeA, b},
		{oppositeB, b},
	}
}

// Does a raycast from a => b, reporting the # of collisions that occur along the way
// ref: https://bryceboe.com/2006/10/23/line-segment-intersection-algorithm/
func Raycast(a, b Vec2, polygon []Vec2) bool {
	for i := range polygon {
		c := polygon[i]
		d := polygon[(i+1)%len(polygon)]

		all := []Vec2{a, b, c, d}

		anySame := false
		for i := range all {
			for j := i + 1; j < len(all); j++ {
				if all[i] == all[j] {
					anySame = true
					break
				}
			}
		}

		if anySame {
			continue
		}

		if IsPointOnLine(a, c, d) || IsPointOnLine(b, c, d) {
			continue
		}

		if doLinesIntersect(a, b, c, d) {
			return true
		}
	}
	return false
}

func IsPointOnLine(point, a, b Vec2) bool {
	if point.x == a.x && point.x == b.x {
		sml, big := MinToMax(a.y, b.y)
		return point.y > sml && point.y < big
	} else if point.y == a.y && point.y == b.y {
		sml, big := MinToMax(a.x, b.x)
		return point.x > sml && point.x < big
	}

	return false
}

func doLinesIntersect(a, b, c, d Vec2) bool {
	return ccw(a, c, d) != ccw(b, c, d) && ccw(a, b, c) != ccw(a, b, d)
}

// Determines if the given points are listed in counter-clockwise order
func ccw(a, b, c Vec2) bool {
	return (c.y-a.y)*(b.x-a.x) > (b.y-a.y)*(c.x-a.x)
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func MinToMax(a, b int) (min, max int) {
	if a > b {
		return b, a
	}
	return a, b
}

func PrintShape(lines [][]Vec2, pointerIndex int) {
	msg := ""
	for i, p := range lines {
		marker := ""
		if i == pointerIndex {
			marker = "<"
		}
		msg += fmt.Sprintf("%s%s\n%s%s\n", p[0].String(), marker, p[1].String(), marker)
	}
	fmt.Println(msg)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
