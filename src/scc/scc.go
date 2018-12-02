package main

import (
	"fmt"
	"os"
	"sort"
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
// Retorna un contador que permite llevar un conteo de los elementos del SCC

func (g *Graph) DFS(n *Node) int {
	n.visited = true

	c := 0

	if len(n.neighbors) != 0 {
		for _, v := range n.neighbors[n.label] {
			if v.visited == false {
				c += g.DFS(v)
			}
		}
	}

	return c + 1
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

// Funcion que se encarga de buscar los componentes fuertemente conectados en el grafo

func (g *Graph) SCC(bytesRead []byte) {
	start := time.Now()

	s := NewStack()

	var n []int
	c := 0

	// Primer dfs
	for label := range g.nodes {
		if g.nodes[label].visited == false {
			g.fillOrder(g.nodes[label], s)
		}
	}

	// Se crea el grafo inverso
	gr := CreateGraph(bytesRead, true)

	// Se realiza el segundo dfs
	for s.Len() > 0 {
		v := (s.Pop()).(int)

		if gr.nodes[v].visited == false {
			c = gr.DFS(gr.nodes[v])
			n = append(n, c)
		}
	}

	// Go organiza los conteos de los elementos de los SCC de mayor a menor
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Print("5 SCC: ")

	// Se imprimen los 5 SCC necesarios

	for i := 0; i < 5; i++ {
		// En caso de que hayan SCC, se imprimen
		if i < len(n) {
			fmt.Print(n[i])
		} else {
			fmt.Print(0) // En caso de que no, se imprime 0
		}
	}

	fmt.Println()

	var elapsed time.Duration
	elapsed = time.Since(start)

	fmt.Printf("SCC time %s\n", elapsed)
}

// Se lee todo el archivo y se carga en una variable

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

// Funcion que devuelve un nodo en caso de que exista en el grafo

func (g *Graph) GetNode(l int) *Node {

	if g.nodes[l] != nil {
		return g.nodes[l]
	}

	return nil
}

// Funcion que crea e inicializa un nodo

func CreateNode(l int) *Node {
	n := new(Node)
	n.label = l
	n.neighbors = make(map[int][]*Node)
	return n
}

// Funcion que crea una arista entre un nodo y otro

func (g *Graph) AddEdge(ini *Node, fin *Node) {
	ini.neighbors[ini.label] = append(ini.neighbors[ini.label], fin)
}

// Funcion que crea el grafo
// r -> true crea el grafo inverso

func CreateGraph(bytes []byte, r bool) *Graph {
	nodes := strings.Fields(string(bytes))

	start := time.Now()

	g := new(Graph)
	g.nodes = make(map[int]*Node)

	for i := 0; i < len(nodes); i += 2 {
		lIni, err := strconv.Atoi(nodes[i])
		lFin, err2 := strconv.Atoi(nodes[i+1])

		if err != nil || err2 != nil {
			fmt.Println("ERROR")
			return nil
		}

		var ini *Node
		var fin *Node

		if g.GetNode(lIni) == nil {
			ini = CreateNode(lIni)
			g.AddNode(ini)
		} else {
			ini = g.GetNode(lIni)
		}
		if g.GetNode(lFin) == nil {
			fin = CreateNode(lFin)
			g.AddNode(fin)
		} else {
			fin = g.GetNode(lFin)
		}

		if r == false {
			g.AddEdge(ini, fin)
		} else {
			g.AddEdge(fin, ini)
		}

		if i%100000 == 0 && !r {
			fmt.Printf("%8d Creating Normal...\n", i)
		} else if i%100000 == 0 && r {
			fmt.Printf("%8d Creating Reverse...\n", i)
		}
	}

	var elapsed time.Duration
	elapsed = time.Since(start)

	fmt.Printf("Creation took %s\n", elapsed)

	return g
}

// Funcion que agrega un nodo ya creado al grafo
func (g *Graph) AddNode(n *Node) {
	g.nodes[n.label] = n
}

func main() {
	name := "SCC.txt"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	// Se lee el archivo y se carga a esta variable
	bytesRead := GetFile(name)
	// Se crea el grafo normal
	gr := CreateGraph(bytesRead, false)
	// Se ejecuta el algoritmo en busca de SCC
	gr.SCC(bytesRead)
}
