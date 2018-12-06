package main

import (
	"fmt"
	"math/rand"
)

type Edge struct {
	source int
	dest   int
}

type Graph struct {
	V int //numero de  vertices
	E int //numero de aristas

	edge []Edge
}

type Subset struct {
	parent int
	rank   int
}

func find(subsets []Subset, i int) int {
	if subsets[i].parent != i {
		subsets[i].parent = find(subsets, subsets[i].parent)
	}

	return subsets[i].parent
}

func Union(subsets []Subset, x int, y int) {
	xroot := find(subsets, x)
	yroot := find(subsets, y)

	if subsets[xroot].rank < subsets[yroot].rank {
		subsets[xroot].parent = yroot
	} else if subsets[xroot].rank > subsets[yroot].rank {
		subsets[yroot].parent = xroot
	} else {
		subsets[yroot].parent = xroot
		subsets[xroot].rank++
	}
}

func createGraph(V int, E int) *Graph {
	g := new(Graph)
	g.V = V
	g.E = E
	g.edge = make([]Edge, E)
	return g
}

func kargerMinCut(g *Graph) int {
	V := g.V
	E := g.E

	edges := g.edge

	subsets := make([]Subset, V)

	for v := 0; v < V; v++ {
		subsets[v].parent = v
		subsets[v].rank = 0
	}

	vertices := V

	for {
		if vertices < 2 {
			break
		}

		i := rand.Intn(E)

		subset1 := find(subsets, edges[i].source)
		subset2 := find(subsets, edges[i].dest)

		if subset1 != subset2 {
			fmt.Printf("Contracting edge %d-%d\n", edges[i].source, edges[i].dest)
			vertices--
			Union(subsets, subset1, subset2)
		}
	}

	cutedges := 0

	for i := 0; i < E; i++ {
		subset1 := find(subsets, edges[i].source)
		subset2 := find(subsets, edges[i].dest)
		if subset1 != subset2 {
			cutedges++
		}
	}

	return cutedges
}

func main() {
	v := 4
	e := 5
	g := createGraph(v, e)

	g.edge[0].source = 0
	g.edge[0].dest = 1

	g.edge[1].source = 0
	g.edge[1].dest = 2

	g.edge[2].source = 0
	g.edge[2].dest = 3

	g.edge[3].source = 1
	g.edge[3].dest = 3

	g.edge[4].source = 2
	g.edge[4].dest = 3

	fmt.Printf("\nKarget's Algorithm: %d", kargerMinCut(g))
}
