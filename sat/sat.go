package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

// Código facilitado por el profesor
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
	//fmt.Println("#", g.nrnodes)
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
		//        s += "\n"
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
		if i%10000 == 0 {
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

	count := 0

	fmt.Printf("%d ", node.label)

	if len(node.neighbors) != 0 {
		for _, v := range node.neighbors {
			if v.visited == false {
				count += g.dfs(v)
			}
		}
	}

	return count + 1
}

//

// Se visita un nodo y se comprueban todos los nodos a los que se puede llegar desde aquí, luego de terminar de un nodo,
// se guarda en el stack
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

	//fmt.Println("PEEK:", (stack.Peek()).(int))
}

// Función que se encarga de procesar e imprimir los SCC resultantes
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

		//fmt.Println("POP: ", v)
		//fmt.Println("NODO: ", gr.getNode(v).label)

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

	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Print("Los 5 SCC más grandes: ")

	for i := 0; i < 5; i++ {
		if i != 4 {
			fmt.Print(n[i], ", ")
		} else {
			fmt.Print(n[i])
		}
	}

	fmt.Println()
}

//

func addEdges(a int, b int, adj [][]int) {
	adj[a] = append(adj[a], b)
}

func addEdgesInverse(a int, b int, adjInv [][]int) {
	adjInv[b] = append(adjInv[b], a)
}

func dfsFirst(u int, visited []bool, adj [][]int, s *Stack) {
	if visited[u] {
		return
	}

	visited[u] = true

	for i := 0; i < len(adj[u]); i++ {
		dfsFirst(adj[u][i], visited, adj, s)
	}

	s.Push(u)
}

func dfsSecond(u int, visitedInv []bool, adjInv [][]int, scc []int, counter int, s *Stack) {
	if visitedInv[u] {
		return
	}

	visitedInv[u] = true

	for i := 0; i < len(adjInv[u]); i++ {
		dfsSecond(adjInv[u][i], visitedInv, adjInv, scc, counter, s)
	}

	scc[u] = counter
}

func sat(n int, lsls [][2]int) {
	adj := [][]int{}
	adjInv := [][]int{}
	visited := make([]bool, n)
	visitedInv := make([]bool, n)
	s := NewStack()

	scc := make([]int, n)
	counter := 1

	a := []int{}
	b := []int{}

	for i := 0; i < len(lsls); i++ {
		a = append(a, lsls[i][0])
		b = append(b, lsls[i][1])
	}

	for i := 0; i < len(lsls); i++ {
		if a[i] > 0 && b[i] > 0 {
			addEdges(a[i]+n, b[i], adj)
			addEdgesInverse(a[i]+n, b[i], adjInv)
			addEdges(b[i]+n, a[i], adj)
			addEdgesInverse(b[i]+n, a[i], adjInv)
		} else if a[i] > 0 && b[i] < 0 {
			addEdges(a[i]+n, n-b[i], adj)
			addEdgesInverse(a[i]+n, n-b[i], adjInv)
			addEdges(-b[i], a[i], adj)
			addEdgesInverse(-b[i], a[i], adjInv)
		} else if a[i] < 0 && b[i] > 0 {
			addEdges(-a[i], b[i], adj)
			addEdgesInverse(-a[i], b[i], adjInv)
			addEdges(b[i]+n, n-a[i], adj)
			addEdgesInverse(b[i]+n, n-a[i], adjInv)
		} else {
			addEdges(-a[i], n-b[i], adj)
			addEdgesInverse(-a[i], n-b[i], adjInv)
			addEdges(-b[i], n-a[i], adj)
			addEdgesInverse(-b[i], n-a[i], adjInv)
		}
	}

	for i := 1; i <= 2*n; i++ {
		if !visited[i] {
			dfsFirst(i, visited, adj, s)
		}
	}

	for s.Len() > 0 {
		n := (s.Pop()).(int)

		if !visitedInv[n] {
			dfsSecond(n, visitedInv, adjInv, scc, counter, s)
			counter = counter + 1
		}
	}

	for i := 1; i <= n; i++ {
		if scc[i] == scc[i+n]
		{
			fmt.Println("Unsastisfiable")
			return
		}
	}

	fmt.Println("Satisfiable")
	return
}

//

func main() {
	start := time.Now()
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//var pg *Graph
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var cantVar int
	var lisnod [][2]int
	var anod [2]int
	i := 0
	var elapsed time.Duration
	for scanner.Scan() {
		lineStr := scanner.Text()
		if i == 0 {
			fmt.Sscanf(lineStr, "%d", &cantVar)
		} else {
			fmt.Sscanf(lineStr, "%d %d", &anod[0], &anod[1])
			lisnod = append(lisnod, anod)
		}
		i++
		if i%10000 == 0 {
			elapsed = time.Since(start)
			fmt.Printf("%8d %v %8d %8d \n", i, elapsed, anod[0], anod[1])
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("Reading took %s\n", elapsed)
	nr := cantVar * 2
	fmt.Printf("Entradas %10d  Nodos %10d\n", len(lisnod), nr)
	//pg = newGraph(nr)
	//pg.creatGraph(lisnod, false, start)
	elapsed = time.Since(start)
	//fmt.Printf("Nodos %10d Creating  %s\n", len(pg.nodes), elapsed)
	//pg.display()
	//pg.dfs(pg.nodes[0])
	//pg.printSCC(lisnod, start)

	sat(cantVar, lisnod)

	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)
}
