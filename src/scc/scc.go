package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Estas funciones permiten el uso de la estructura de datos "Stack" en Go
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

// Estructura que define un nodo

type Node struct {
	label     int             // identificador
	visited   bool            // ha sido visitado?
	neighbors map[int][]*Node // diccionario que contiene los vecinos de este nodo
}

// Estructura que define un grafo

type Graph struct {
	nodes map[int]*Node // diccionario que contiene todos los nodos del grafo
}

// Funcion que realiza el recorrido de profundidad
// Retorna un contador que permite llevar un conteo de los elementos del SCC

func (g *Graph) dfs(node *Node) int {
	node.visited = true

	count := 0

	if len(node.neighbors) != 0 {
		for _, v := range node.neighbors[node.label] {
			if v.visited == false {
				count += g.dfs(v)
			}
		}
	}

	return count + 1
}

// Se hace un recorrido de profundidad, agregando cada nodo visitado al stack

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

// Funcion que se encarga de buscar los componentes fuertemente conectados en el grafo

func (g *Graph) printSCC(bytesRead []byte) {
	start := time.Now()

	stack := NewStack()

	var n []int
	count := 0

	// Primer dfs
	for label := range g.nodes {
		if g.nodes[label].visited == false {
			g.fillOrder(g.nodes[label], stack)
		}
	}

	// Se crea el grafo inverso
	gr := CreateGraph(bytesRead, true)

	// Se realiza el segundo dfs
	for stack.Len() > 0 {
		v := (stack.Pop()).(int)

		if gr.nodes[v].visited == false {
			count = gr.dfs(gr.nodes[v])
			n = append(n, count)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Print("5 SCC: ")

	for i := 0; i < 5; i++ {
		if i < len(n) {
			if i != 4 {
				fmt.Print(n[i], ", ")
			} else {
				fmt.Print(n[i])
			}
		} else {
			if i != 4 {
				fmt.Print(0, ", ")
			} else {
				fmt.Print(0)
			}
		}
	}

	fmt.Println()

	var elapsed time.Duration
	elapsed = time.Since(start)

	fmt.Printf("SCC time %s\n", elapsed)
}

// Se lee todo el archivo y se carga en una variable

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

	fmt.Printf("Reading time %s\n", elapsed)

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

	gr.printSCC(bytesRead)

}
