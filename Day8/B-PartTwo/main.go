package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Vec3 struct {
	x, y, z int
}

type Connection struct {
	a, b int
	dist int
}

func main() {
	path := "../input.txt"

	buf, err := os.Open(path)
	handleErr(err)
	defer buf.Close()
	scanner := bufio.NewScanner(buf)

	inputs := make([]Vec3, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		handleErr(err)
		y, err := strconv.Atoi(parts[1])
		handleErr(err)
		z, err := strconv.Atoi(parts[2])
		handleErr(err)

		inputs = append(inputs, Vec3{x: x, y: y, z: z})
	}

	// Make all connections
	connections := make([]Connection, 0)
	for i := 0; i < len(inputs)-1; i++ {
		a := inputs[i]
		for j := i + 1; j < len(inputs); j++ {
			b := inputs[j]
			dist := a.DstSqrd(b)
			connections = append(connections, Connection{a: i, b: j, dist: dist})
		}
	}

	// Order by closest connections
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].dist < connections[j].dist
	})

	// Take the desired number of closest connections
	desiredConnections := 1000
	formedConnections := connections[:desiredConnections]

	// Create adjacency map
	adj := make(map[int][]int, 0)
	for _, c := range formedConnections {
		adj[c.a] = append(adj[c.a], c.b)
		adj[c.b] = append(adj[c.b], c.a)
	}

	res := countConnectedRegionSizes(adj)
	slices.Sort(res)
	slices.Reverse(res)

	sum := 1
	for i := range 3 {
		sum *= res[i]
	}

	fmt.Println(sum)
}

func countConnectedRegionSizes(adj map[int][]int) []int {
	visited := make(map[int]bool, 0)

	counts := make([]int, 0)
	// Go over each node and its connections. If it's already been visited
	// we have already counted the node as part of another region
	for k := range adj {
		if !visited[k] {
			counts = append(counts, bfsConnected(adj, k, visited))
		}
	}
	return counts
}

// Breadth first search through the adjacencies. Doesn't return the visited nodes
// just counts the nodes it finds
func bfsConnected(adj map[int][]int, start int, visited map[int]bool) int {
	// A faux queue
	q := make([]int, 1)
	q[0] = start
	visited[start] = true

	count := 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		count++

		for _, n := range adj[curr] {
			if !visited[n] {
				visited[n] = true
				q = append(q, n)
			}
		}
	}
	return count
}

// Calculates distance squared from a to b. No need to take square root since
// we're looking at relative distances
func (a *Vec3) DstSqrd(b Vec3) int {
	return Squared(a.x-b.x) + Squared(a.y-b.y) + Squared(a.z-b.z)
}

func Squared(x int) int {
	return x * x
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
