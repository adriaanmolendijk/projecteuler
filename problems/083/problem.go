package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Graph implementation from following source
// https://tobin.cc/blog/graph/

type Graph struct {
	NumNodes int
	Edges    [][]Edge
}

type Edge struct {
	From   int
	To     int
	Weight int
}

func NewGraph(n int) *Graph {
	return &Graph{
		NumNodes: n,
		Edges:    make([][]Edge, n),
	}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.Edges[u] = append(g.Edges[u], Edge{From: u, To: v, Weight: w})

	// For undirected graph add edge from v to u.
	// g.Edges[v] = append(g.Edges[v], Edge{From: v, To: u, Weight: w})
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	n := 80
	grid, row := make([][]int, n), 1
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ",")

		grid[row-1] = make([]int, n)
		for col := 1; col <= n; col++ {
			num, _ := strconv.ParseInt(numbers[col-1], 10, 32)
			grid[row-1][col-1] = int(num)
		}
		row++
	}

	g := NewGraph(n * n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			vertex := i - 1 + (j-1)*n

			// first row has no edges
			if i > 1 {
				vertexAbove := vertex - 1
				g.AddEdge(vertex, vertexAbove, grid[i-2][j-1])
			}

			// last row has no edges
			if i < n {
				vertexBelow := vertex + 1
				g.AddEdge(vertex, vertexBelow, grid[i][j-1])
			}

			// first col has no edges
			if j > 1 {
				vertexLeft := vertex - n
				g.AddEdge(vertex, vertexLeft, grid[i-1][j-2])
			}

			// last col has no edges
			if j < n {
				vertexRight := vertex + n
				g.AddEdge(vertex, vertexRight, grid[i-1][j])
			}
		}
	}
	dist := dijkstrasAlgorithm(*g, 0)
	fmt.Println(dist[n*n-1] + grid[0][0])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func dijkstrasAlgorithm(g Graph, source int) []int {
	n := g.NumNodes
	dist := make([]int, n)
	prev := make([]int, n)
	Q := map[int]bool{}
	for i := 0; i < n; i++ {
		dist[i] = 1 << 31 // Set to infinity
		prev[i] = -1
		Q[i] = true
	}
	dist[source] = 0

	for len(Q) > 0 {
		// Select vertex in unexplored with minimum cost
		u := findClosestQ(dist, Q)
		delete(Q, u)

		edges := g.Edges[u]
		for _, edge := range edges {
			alt := dist[edge.From] + edge.Weight
			if alt < dist[edge.To] {
				dist[edge.To] = alt
				prev[edge.To] = edge.From
			}
		}
	}

	return dist
}

func findClosestQ(dist []int, Q map[int]bool) int {
	minVertex := -1
	minCost := 1 << 31
	for vertex := range Q {
		if dist[vertex] < minCost {
			minCost = dist[vertex]
			minVertex = vertex
		}
	}

	if minVertex == -1 {
		return 0
	}
	return minVertex
}
