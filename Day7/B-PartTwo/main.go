package main

import (
	"bufio"
	"fmt"
	"os"
)

// Represents either a splitter or the starting position
type Node struct {
	// X & Y coords
	x, y int
	// Pointers to the next node this hits. Nil if it reaches the end
	leftConnection, rightConnection *Node
	// Cache possible paths once counted
	possiblePaths int
}

func main() {
	path := "../input.txt"

	buf, err := os.Open(path)
	handleErr(err)
	defer buf.Close()
	scanner := bufio.NewScanner(buf)

	// 2D array with top level array being a column, each nested array is a row.
	// This will allow us to easily find the next node down in a column or determine if
	// we're at the last one
	nodeColumns := make([][]*Node, 0)

	y := -1
	var entryNode *Node
	for scanner.Scan() {
		line := scanner.Text()
		y++

		for x, char := range line {
			if x >= len(nodeColumns) {
				nodeColumns = append(nodeColumns, make([]*Node, 0))
			}

			if char != 'S' && char != '^' {
				continue
			}
			node := Node{x: x, y: y, possiblePaths: -1}
			nodeColumns[x] = append(nodeColumns[x], &node)

			if char == 'S' {
				entryNode = &node
			}
		}
	}

	formConnections(nodeColumns)
	result := countPossiblePaths(entryNode)

	fmt.Println("There are", result, "possilbe paths")
}

// Given x,y coordinates, finds the next node down or nil if it reaches the bottom
func findNextNode(nodes [][]*Node, x, y int) *Node {
	if x >= len(nodes) {
		return nil
	}
	column := nodes[x]

	for _, node := range column {
		if node.y > y {
			return node
		}
	}
	return nil
}

// Finds the next node down on the left & right side of each node
func formConnections(nodes [][]*Node) {
	for x, column := range nodes {
		for _, node := range column {
			node.leftConnection = findNextNode(nodes, x-1, node.y)
			node.rightConnection = findNextNode(nodes, x+1, node.y)
		}
	}
}

// Recursively counts how many paths proceed the given node
func countPossiblePaths(node *Node) int {
	paths := 0

	if node.possiblePaths != -1 {
		return node.possiblePaths
	}

	if node.leftConnection == nil {
		paths += 1
	} else {
		paths += countPossiblePaths(node.leftConnection)
	}

	if node.rightConnection == nil {
		paths += 1
	} else {
		paths += countPossiblePaths(node.rightConnection)
	}

	node.possiblePaths = paths
	return paths
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
