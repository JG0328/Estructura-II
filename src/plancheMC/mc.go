package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Everything related to graphs

// Graph maintains slices of graph vertices and edges
type Graph struct {
	Vertices int
	Edges    []edge
}

type edge struct {
	tail int
	head int
}

// New returns pointer to a new graph struct
func New() *Graph {
	return &Graph{}
}

// AddEdge appends new edge struct to a graph's edges slice
func (g *Graph) AddEdge(tail, head int) {
	e := edge{tail, head}
	g.Edges = append(g.Edges, e)
}

// AddVertex increments vertices count of a Graph struct
func (g *Graph) AddVertex(label int) {
	g.Vertices++
}

// ContractEdge will collapse an edge and merge vertex elements
func (g *Graph) ContractEdge(i int) {
	e := g.Edges[i]

	// update edges to remove e
	g.Edges = append(g.Edges[:i], g.Edges[i+1:]...)

	// loop to remove e.head from edges slice
	for j, edge := range g.Edges {
		if edge.head == e.head {
			g.Edges[j].head = e.tail
		} else if edge.tail == e.head {
			g.Edges[j].tail = e.tail
		}
	}

	// remove self loops
	temp := []edge{}
	for _, edge := range g.Edges {
		if edge.head != edge.tail {
			temp = append(temp, edge)
		}
	}
	g.Edges = temp

	g.Vertices--
}

// Functions

// Karger takes a pointer to Graph struct as argument
// returns min cut of graph based on Kerger alg, nondeterministic
// set listedTwice to true if edges are represented twice in graph
func Karger(g *Graph, listedTwice bool) int {
	min := len(g.Edges)
	for i := 0; i < 1000; i++ {
		if m := iterate(g, listedTwice); m < min {
			min = m
		}
	}

	return min
}

func iterate(g *Graph, listedTwice bool) int {
	for g.Vertices > 2 {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		i := r1.Intn(len(g.Edges))

		g.ContractEdge(i)
	}

	if listedTwice {
		return len(g.Edges) / 2
	}
	return len(g.Edges)
}

// Read file and create a graph
func graphFromTSV(name string) *Graph {
	input, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal("READING FILE ERROR")
	}

	// initialize new graph struct
	g := New()

	// split by line
	s := strings.Split(string(input), "\n")
	for _, line := range s {
		items := strings.Split(line, "\t")
		tail, _ := strconv.Atoi(items[0])
		for i := 1; i < len(items); i++ {
			head, _ := strconv.Atoi(items[i])
			g.AddEdge(tail, head)
		}
		g.Vertices++
	}

	return g
}

//

func main() {
	var name string

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	g := graphFromTSV(name)

	var min int

	for i := 0; i < 100; i++ {
		a := Karger(g, true)

		if i == 0 || a < min {
			min = a
		}
	}

	fmt.Println("Result -> ", min)
}
