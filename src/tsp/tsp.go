package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	label     int
	visited   bool
	neighbors map[int]*Edge
	position  Position
}

type Graph struct {
	nodes map[int]*Node
}

type Position struct {
	posX float64
	posY float64
}

type Edge struct {
	node   *Node
	weight float64
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

func (g *Graph) GetNode(label int) *Node {

	if g.nodes[label] != nil {
		return g.nodes[label]
	}

	return nil
}

func CreateNode(label int) *Node {
	n := new(Node)
	n.label = label
	n.neighbors = make(map[int]*Edge)
	return n
}

func (g *Graph) AddEdge(nini *Node, nfin *Node) {
	//nini.neighbors[nini.label] = append(nini.neighbors[nini.label], nfin)
}

func CreateEdge(node *Node, weight float64) *Edge {
	e := new(Edge)
	e.node = node
	e.weight = weight

	return e
}

func SetDistance(nini *Node, nfin *Node) {
	x := math.Pow(nfin.position.posX-nini.position.posX, 2)
	y := math.Pow(nfin.position.posY-nini.position.posY, 2)
	z := math.Sqrt(x + y)

	nini.neighbors[nfin.label] = CreateEdge(nfin, z)
	nfin.neighbors[nini.label] = CreateEdge(nini, z)
}

func CreateGraph(bytesRead []byte, rev bool) *Graph {
	nodes := strings.Fields(string(bytesRead))

	start := time.Now()

	g := new(Graph)
	g.nodes = make(map[int]*Node)

	for i := 0; i < len(nodes); i += 2 {
		labelIni, err := strconv.Atoi(nodes[i])
		labelFin, err2 := strconv.Atoi(nodes[i+1])

		if err != nil || err2 != nil {
			fmt.Println("ERROR CREATING THE GRAPH")
			return nil
		}

		var nini *Node
		var nfin *Node

		if g.GetNode(labelIni) == nil {
			nini = CreateNode(labelIni)
			g.AddNode(nini)
		} else {
			nini = g.GetNode(labelIni)
		}
		if g.GetNode(labelFin) == nil {
			nfin = CreateNode(labelFin)
			g.AddNode(nfin)
		} else {
			nfin = g.GetNode(labelFin)
		}

		if rev == false {
			g.AddEdge(nini, nfin)
		} else {
			g.AddEdge(nfin, nini)
		}

		if i%100000 == 0 && !rev {
			fmt.Printf("%8d - Creating...\n", len(g.nodes))
		} else if i%100000 == 0 && rev {
			fmt.Printf("%8d - Creating Reverse...\n", len(g.nodes))
		}
	}

	var elapsed time.Duration
	elapsed = time.Since(start)

	fmt.Printf("Creation took %s\n", elapsed)

	return g
}

func (g *Graph) AddNode(node *Node) {
	g.nodes[node.label] = node
}

func main() {
	name := "SCC.txt"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	bytesRead := ReadFile(name)

	if bytesRead == nil {
		return
	}

	gr := CreateGraph(bytesRead, false)

	if gr == nil {
		return
	}

}
