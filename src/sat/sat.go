package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

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

//

// Funcion que realiza el recorrido de profundidad
// Va agregando los nodos que visita a un diccionario

func (g *Graph) dfs(n *Node) int {
	n.visited = true

	if len(n.neighbors) != 0 {
		for _, v := range n.neighbors[n.label] {
			if v.visited == false {
				sccDic[g.dfs(v)] = v
			}
		}
	}

	return n.label
}

// Se hace un recorrido de profundidad, agregando cada nodo visitado al stack

func (g *Graph) fillOrder(n *Node, s *Stack) {
	n.visited = true

	if len(n.neighbors) != 0 {
		for _, v := range n.neighbors[n.label] {
			if v.visited == false {
				g.fillOrder(v, s)
			}
		}
	}

	s.Push(n.label)
}

// Funcion que se encarga de procesar los SCC y determinar si 2sat tiene solucion o no

func (g *Graph) GetSCC(bytes []byte) {
	solution = 1 // inicialmente, el problema tiene solucion

	sccDic = make(map[int]*Node) // inicializo el diccionario de SCC

	start := time.Now()

	s := NewStack()

	// Se colocan los nodos en el stack
	for label := range g.nodes {
		if g.nodes[label].visited == false {
			g.fillOrder(g.nodes[label], s)
		}
	}

	// Se crea el grafo inverso
	gr := CreateGraph(bytes, true)

	for s.Len() > 0 {
		v := (s.Pop()).(int)

		if gr.nodes[v].visited == false {
			sccDic[gr.dfs(gr.nodes[v])] = gr.nodes[v]

			// x y ~n existen en el mismo SCC? Entonces no tiene solucion
			if sccDic[v] != nil && sccDic[v*-1] != nil {
				solution = 0
			}

			sccDic = nil                 // Se limpia el diccionario
			sccDic = make(map[int]*Node) // Se inicializa de nuevo
		}
	}

	if solution == 1 {
		fmt.Println("Se puede resolver")
	} else {
		fmt.Println("No se puede resolver")
	}

	var elapsed time.Duration
	elapsed = time.Since(start)

	fmt.Printf("SCC time %s\n", elapsed)
}

// Funcion que se encarga de leer el archivo

func GetFile(n string) []byte {
	start := time.Now()
	file, err := os.Open(n)

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

	file.Read(buffer)

	var elapsed time.Duration
	elapsed = time.Since(start)

	fmt.Printf("Reading time %s\n", elapsed)

	return buffer
}

// Funcion que devuelve un nodo si existe un nodo con este id en el grafo

func (g *Graph) GetNode(l int) *Node {

	if g.nodes[l] != nil {
		return g.nodes[l]
	}

	return nil
}

// Funcion que inicializa un nodo

func CreateNode(label int) *Node {
	n := new(Node)
	n.label = label
	n.neighbors = make(map[int][]*Node)
	return n
}

// Funcion que agrega un nodo al grafo

func (g *Graph) AddNode(node *Node) {
	g.nodes[node.label] = node
}

// Funcion que agrega una arista entre un nodo y otro

func (g *Graph) AddEdge(nini *Node, nfin *Node) {
	nini.neighbors[nini.label] = append(nini.neighbors[nini.label], nfin)
}

// Funcion que se encarga de crear un grafo
// r -> true crea el grafo inverso

func CreateGraph(b []byte, r bool) *Graph {
	nodes := strings.Fields(string(b))

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

		if r == false {
			g.AddEdge(nini, nfin)
		} else {
			g.AddEdge(nfin, nini)
		}

		if i%100000 == 0 && !r {
			fmt.Printf("%8d - Creating Normal...\n", len(g.nodes))
		} else if i%100000 == 0 && r {
			fmt.Printf("%8d - Creating Reverse...\n", len(g.nodes))
		}
	}

	var elapsed time.Duration
	elapsed = time.Since(start)

	fmt.Printf("Creation time %s\n", elapsed)

	return g
}

var sccDic map[int]*Node // diccionario temporal de SCC
var solution int         // variable que determina si el problema tiene solucion o no

func main() {
	var n string
	if len(os.Args) > 1 {
		n = os.Args[1]
	}

	// Leo el archivo y obtengo los datos
	bytes := GetFile(n)

	// Creo el grafo
	gr := CreateGraph(bytes, false)

	// Analizo los SCC para saber si el problema tiene o no solucion
	gr.GetSCC(bytes)
}
