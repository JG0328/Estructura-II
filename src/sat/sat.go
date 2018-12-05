package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Implementación del Stack
type (
	Stack struct {
		top    *stackNode
		length int
	}
	stackNode struct {
		value interface{}
		prev  *stackNode
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) Len() int {
	return this.length
}

func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *Stack) Push(value interface{}) {
	n := &stackNode{value, this.top}
	this.top = n
	this.length++
}

//

type Node struct {
	label     int
	visited   bool
	neighbors map[int][]*Node
}

type Graph struct {
	nodes map[int]*Node
}

// El DFS servira para verificar si x y ~x se encuentran son SCC

func (g *Graph) dfs(node *Node, sccMap map[int]*Node) {
	node.visited = true

	sccMap[node.label] = node

	if len(node.neighbors) != 0 {
		for _, v := range node.neighbors[node.label] {
			if v.visited == false {
				g.dfs(v, sccMap)
			}
		}
	}

	return
}

// Se visita un nodo y se comprueban todos los nodos a los que se puede llegar desde aquí, luego de terminar de un nodo,
// se guarda en el stack

func (g *Graph) fillOrder(node *Node, stack *Stack) {
	node.visited = true

	if len(node.neighbors) != 0 {
		for _, v := range node.neighbors[node.label] {
			if v.visited == false {
				g.fillOrder(v, stack)
			}
		}
	}

	stack.Push(node.label)
}

// Función que se encarga de procesar el SCC e imprimir 0 o 1 si tiene solucion

var sccMap map[int]*Node
var sat int

func (g *Graph) GetSCC(bytesRead []byte) {
	sat = 1

	sccMap = make(map[int]*Node)

	start := time.Now()

	stack := NewStack()

	// Se colocan los nodos en el stack
	for label := range g.nodes {
		if g.nodes[label].visited == false {
			g.fillOrder(g.nodes[label], stack)
		}
	}

	gr := CreateGraph(bytesRead, true)

	for stack.Len() > 0 {
		v := (stack.Pop()).(int)

		if gr.nodes[v].visited == false {
			gr.dfs(gr.nodes[v], sccMap)

			if sccMap[v] != nil && sccMap[v*-1] != nil {
				sat = 0
			}

			/*
				if len(sccMap) > 1 {
					fmt.Println("SCC -> ", sccMap)
				}
			*/

			sccMap = nil
			sccMap = make(map[int]*Node)
		}
	}

	if sat == 1 {
		fmt.Println("Tiene solucion")
	} else {
		fmt.Println("No tiene solucion")
	}

	var elapsed time.Duration
	elapsed = time.Since(start)

	fmt.Printf("SCC took %s\n", elapsed)
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
	n.neighbors = make(map[int][]*Node)
	return n
}

func (g *Graph) AddEdge(nini *Node, nfin *Node) {
	nini.neighbors[nini.label] = append(nini.neighbors[nini.label], nfin)
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
	var name string
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

	gr.GetSCC(bytesRead)

}
