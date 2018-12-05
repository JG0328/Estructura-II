package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

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

type Node struct {
	label     int
	visited   bool
	neighbors []*Node
}

type Graph struct {
	nrnodes int
	nodes   []*Node
}

func (g *Graph) AddNode(n *Node) {
	g.nodes[g.nrnodes] = n
	g.nrnodes++
}

func (n *Node) AddEdge(nd *Node) {
	n.neighbors = append(n.neighbors, nd)
}

func (g *Graph) getNode(n int) *Node {
	for i := 0; i < g.nrnodes; i++ {
		pnd := g.nodes[i]
		if pnd.label == n {
			return pnd
		}
	}
	return nil
}

func (g *Graph) display() {
	s := ""
	for i := 0; i < g.nrnodes; i++ {
		s += fmt.Sprintf("%8d --> ", g.nodes[i].label)
		near := g.nodes[i]
		for j := 0; j < len(near.neighbors); j++ {
			s += fmt.Sprintf("%8d ", near.neighbors[j].label)
		}
		fmt.Println(s)
		s = ""
	}
}

func newGraph(n int) (pg *Graph) {
	pg = new(Graph)
	pg.nodes = make([]*Node, n)
	return
}

func (pg *Graph) creatGraph(lsls [][2]int, rev bool, start time.Time) {
	var i int
	for _, lis := range lsls {
		var nini, nfin int
		if rev {
			nini = lis[1]
			nfin = lis[0]
		} else {
			nini = lis[0]
			nfin = lis[1]
		}
		pini := pg.getNode(nini)
		if pini == nil {
			pini = new(Node)
			pini.label = nini
			pg.AddNode(pini)
		}
		nd := pg.getNode(nfin)
		if nd == nil {
			nd = new(Node)
			nd.label = nfin
			pg.AddNode(nd)
		}
		pini.AddEdge(nd)
		i++
		if i%100000 == 0 {
			elapsed := time.Since(start)
			fmt.Printf("%8d %v %8d %8d \n", i, elapsed, nini, nfin)
		}
	}
}

func getMax(lis2 [][2]int) (n int) {
	for _, lis := range lis2 {
		if n < lis[0] {
			n = lis[0]
		}
		if n < lis[1] {
			n = lis[1]
		}
	}
	return n
}

func (g *Graph) dfs(node *Node) int {
	node.visited = true

	cont := 0

	fmt.Printf("%d ", node.label)

	if len(node.neighbors) != 0 {
		for _, v := range node.neighbors {
			if v.visited == false {
				cont += g.dfs(v)
			}
		}
	}

	return cont + 1
}

// Visita los nodos y los guarda en el stack
func (g *Graph) fillOrder(node *Node, stack *Stack) {
	node.visited = true

	if len(node.neighbors) != 0 {
		for _, v := range node.neighbors {
			if v.visited == false {
				g.fillOrder(v, stack)
			}
		}
	}

	stack.Push(node.label - 1)
}

// Se imprimen los SCC
func (g *Graph) printSCC(lisnod [][2]int, start time.Time) {
	stack := NewStack()

	n := make([]int, 5)
	index := 0
	count := 0

	for i := 0; i < len(g.nodes); i++ {
		if g.getNode(i+1).visited == false {
			g.fillOrder(g.getNode(i+1), stack)
		}
	}

	// Se crea el inverso
	gr := newGraph(len(g.nodes))
	gr.creatGraph(lisnod, true, start)
	//

	for stack.Len() > 0 {
		v := (stack.Pop()).(int) + 1

		if gr.getNode(v).visited == false {
			fmt.Print("SCC: ")
			count = gr.dfs(gr.getNode(v))
			if index < 5 {
				n[index] = count
				index = index + 1
			}
			fmt.Println()
		}
	}

	// Se ordenan los scc de mayor a menos
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Print("SCC -> ")

	for i := 0; i < 5; i++ {
		fmt.Print(n[i], " ")
	}

	fmt.Println()
}

func main() {
	start := time.Now()
	name := "SCC.txt"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var pg *Graph
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var lisnod [][2]int
	var anod [2]int
	var i int
	var elapsed time.Duration
	for scanner.Scan() {
		lineStr := scanner.Text()
		fmt.Sscanf(lineStr, "%d %d", &anod[0], &anod[1])
		lisnod = append(lisnod, anod)
		i++
		if i%100000 == 0 {
			elapsed = time.Since(start)
			fmt.Printf("%8d %v %8d %8d \n", i, elapsed, anod[0], anod[1])
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("Reading took %s\n", elapsed)
	nr := getMax(lisnod)
	fmt.Printf("Entradas %10d  Nodos %10d\n", len(lisnod), nr)
	pg = newGraph(nr)
	pg.creatGraph(lisnod, false, start)
	elapsed = time.Since(start)
	fmt.Printf("Nodos %10d Creating  %s\n", len(pg.nodes), elapsed)
	pg.display()
	pg.printSCC(lisnod, start)
	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)
}
