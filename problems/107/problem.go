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

	n, v := 40, 0
	g := NewGraph(n)
	for scanner.Scan() {
		v++
		line := scanner.Text()
		numbers := strings.Split(line, ",")
		for w := 1; w <= n; w++ {
			num := numbers[w-1]
			if num != "-" {
				weight, _ := strconv.ParseInt(num, 10, 64)
				g.AddEdge(v-1, w-1, int(weight))
			}
		}
	}

	totalWeight := 0
	for _, edges := range g.Edges {
		for _, edge := range edges {
			totalWeight += edge.Weight
		}
	}
	fmt.Println(totalWeight/2 - primsAlgorithm(*g))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// https://en.wikipedia.org/wiki/Prim%27s_algorithm
func primsAlgorithm(g Graph) int {
	n := g.NumNodes
	cheapestCost := make([]int, n)
	cheapestEdge := make([]int, n)
	explored := make(map[int]bool)
	unexplored := make(map[int]bool)
	for i := 0; i < n; i++ {
		cheapestCost[i] = 1 << 31 // Set to infinity
		cheapestEdge[i] = -1
		unexplored[i] = true
	}

	for len(unexplored) > 0 {
		// Select vertex in unexplored with minimum cost
		currentVertex := findCheapestVertex(cheapestCost, unexplored)
		delete(unexplored, currentVertex)
		explored[currentVertex] = true

		edges := g.Edges[currentVertex]
		for _, edge := range edges {
			if unexplored[edge.To] && edge.Weight < cheapestCost[edge.To] {
				cheapestCost[edge.To] = edge.Weight
				cheapestEdge[edge.To] = currentVertex
			}
		}
	}

	optimalWeight := 0
	resultEdges := []Edge{}
	for i := 0; i < n; i++ {
		if cheapestEdge[i] != -1 {
			resultEdges = append(resultEdges, Edge{From: cheapestEdge[i], To: i, Weight: cheapestCost[i]})
			optimalWeight += cheapestCost[i]
		}
	}
	return optimalWeight
}

func findCheapestVertex(cheapestCost []int, unexplored map[int]bool) int {
	minVertex := -1
	minCost := 1 << 31
	for vertex := range unexplored {
		if cheapestCost[vertex] < minCost {
			minCost = cheapestCost[vertex]
			minVertex = vertex
		}
	}

	if minVertex == -1 {
		return 0
	}
	return minVertex
}
