package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	label     int
	visited   bool
	neighbors map[int]*Node
	distances map[*Node]float64
	position  Position
}

type Graph struct {
	nodes map[int]*Node
}

type Position struct {
	posX float64
	posY float64
}

func ReadFile(name string) []byte {
	start := time.Now()
	file, err := os.Open(name)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var elapsed time.Duration
	elapsed = time.Since(start)

	fmt.Println("bytes read: ", bytesread)

	fmt.Printf("Reading took %s\n", elapsed)

	return buffer
}

func CreateNode(label int) *Node {
	n := new(Node)
	n.label = label
	n.neighbors = make(map[int]*Node)
	n.distances = make(map[*Node]float64)
	return n
}

func (g *Graph) SetDistance(node *Node) {
	node.neighbors = g.nodes

	for i := 0; i < len(node.neighbors); i++ {
		if (i + 1) <= node.label {
			delete(node.distances, node.neighbors[i+1])
			//fmt.Print("Node ", node.label, " deleted... ", node.neighbors[i+1].label)
		}
	}

	for id := range node.neighbors {
		if node.distances[node] == 0 {
			x := math.Pow(node.position.posX-node.neighbors[id].position.posX, 2)
			y := math.Pow(node.position.posY-node.neighbors[id].position.posY, 2)
			z := math.Sqrt(x + y)
			node.distances[node.neighbors[id]] = z
			node.neighbors[id].distances[node] = z
		}
	}
}

func (g *Graph) CreateConnections() {
	start := time.Now()

	for i := 0; i < len(g.nodes); i++ {
		g.SetDistance(g.nodes[i+1])

		if i%1000 == 0 {
			var elapsed time.Duration
			elapsed = time.Since(start)
			fmt.Printf("%8d - Connections - %s\n", i, elapsed)
		}
	}
}

func CreateGraph(bytesRead []byte, rev bool) *Graph {
	data := strings.Fields(string(bytesRead))

	start := time.Now()

	g := new(Graph)
	g.nodes = make(map[int]*Node)

	for i := 0; i < len(data); i += 3 {
		label, err := strconv.Atoi(data[i])

		if err != nil {
			fmt.Println("ERROR CREATING THE GRAPH")
			return nil
		}

		var node *Node

		node = CreateNode(label)

		x, errX := strconv.ParseFloat(data[i+1], 64)
		y, errY := strconv.ParseFloat(data[i+2], 64)

		if errX != nil || errY != nil {
			fmt.Println("ERROR CREATING THE GRAPH")
			return nil
		}

		node.position.posX = x
		node.position.posY = y

		g.AddNode(node)

		if node.label%1000 == 0 {
			fmt.Printf("%8d - Creating...\n", node.label)
		}
	}

	var elapsed time.Duration
	elapsed = time.Since(start)

	fmt.Printf("Creation took %s\n", elapsed)

	fmt.Print("Graph nodes: ", len(g.nodes), "\n")

	return g
}

func (g *Graph) AddNode(node *Node) {
	g.nodes[node.label] = node
}

func (g *Graph) TSP(source *Node) float64 {
	start := time.Now()

	var nodesApart []int

	for i := 0; i < len(g.nodes); i++ {
		if source != g.nodes[i+1] {
			nodesApart = append(nodesApart, i+1)
		}
	}

	var minPath float64
	minPath = -1

	for {
		// Code

		var currentPathWeight float64

		k := source.label

		for i := 0; i < len(nodesApart); i++ {
			currentPathWeight += g.nodes[k].distances[g.nodes[nodesApart[i]]]
			k = nodesApart[i]
		}

		currentPathWeight += g.nodes[k].distances[g.nodes[source.label]]

		minPath = math.Min(minPath, currentPathWeight)

		// Condition
		if !NextPermutation(sort.IntSlice(nodesApart)) {
			break
		}
	}

	var elapsed time.Duration
	elapsed = time.Since(start)

	fmt.Printf("TSP took %s\n", elapsed)

	return minPath
}

// Next permutation
func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func main() {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	bytesRead := ReadFile(name)

	if bytesRead == nil {
		return
	}

	gr := CreateGraph(bytesRead, false)
	gr.CreateConnections()

	fmt.Printf("Resultado: %.2f\n", gr.TSP(gr.nodes[1]))

	if gr == nil {
		return
	}

}
