package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

/*
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
*/

// Código facilitado por el profesor
/*
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
*/

/*
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
*/

// Función que se encarga de procesar e imprimir los SCC resultantes
/*
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
*/

/*
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
	//pg.dfs(pg.nodes[0])
	pg.printSCC(lisnod, start)
	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)
}
*/

type twoInt [2]int

type Node struct {
	label     int // nombre del nodo
	visited   int // 0 not visited, -1 gray, 1 visited
	previo    int // indice a nodo anterior en recorrido
	neighbors []int
}

type Graph struct {
	nodes []*Node
}

var (
	t     int
	pila  []int
	arbol []int
)

func (g *Graph) getNode(n int) (found bool, j int) {
	for i := 0; i < len(g.nodes); i++ {
		pnd := g.nodes[i]
		if pnd.label == n {
			found = true
			j = i
			return
		}
	}
	return
}

func (g *Graph) display() {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		near := g.nodes[i]
		s += fmt.Sprintf(" %8d --> ", near.label)
		for _, j := range near.neighbors {
			s += fmt.Sprintf(" %8d| ", g.nodes[j].label)
		}
		fmt.Println(s)
		s = ""
	}
}

func (pg *Graph) creatGraph(lsls []twoInt, rev bool, start time.Time) {
	var i, pos, posf int
	var node, ndf *Node
	var halla bool
	for _, lis := range lsls {
		var nini, nfin int
		if rev {
			nini = lis[1]
			nfin = lis[0]
		} else {
			nini = lis[0]
			nfin = lis[1]
		}
		halla, pos = pg.getNode(nini)
		if halla {
			node = pg.nodes[pos]
		} else {
			node = new(Node)
			node.label = nini
			pg.nodes = append(pg.nodes, node)
			pos = len(pg.nodes) - 1
		}
		halla, posf = pg.getNode(nfin)
		if halla {
			ndf = pg.nodes[posf]
		} else {
			ndf = new(Node)
			ndf.label = nfin
			pg.nodes = append(pg.nodes, ndf)
			posf = len(pg.nodes) - 1
		}
		node.neighbors = append(node.neighbors, posf)
		i++
		if i%100000 == 0 {
			elapsed := time.Since(start)
			fmt.Printf("%8d %v %8d %8d \n", i, elapsed, nini, nfin)
		}
	}
}

func getMax(lis2 []twoInt) (n int) {
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

func (g *Graph) dfs(rev bool) {
	t = -1
	if !rev {
		for i, u := range g.nodes {
			if u.visited == 0 {
				u.previo = -1
				//                  arbol = append(arbol,i)
				dfsVisit(g, i, u, rev)
			}
		}
	} else {
		for j := len(pila) - 1; j >= 0; j-- {
			i := pila[j]
			u := g.nodes[i]
			if u.visited == 0 {
				u.previo = -1
				arbol = append(arbol, i)
				// fmt.Println("dfs ", g.nodes[i].label)
				dfsVisit(g, i, u, rev)
			}
		}
	}
}

func dfsVisit(g *Graph, i int, u *Node, rev bool) {
	u.visited = -1
	for _, k := range u.neighbors {
		v := g.nodes[k]
		if v.visited == 0 {
			v.previo = i
			if rev {
				arbol[len(arbol)-1] = k
			}
			dfsVisit(g, k, v, rev)
		}
	}
	if !rev {
		t++
		pila = append(pila, i)
	}
	u.visited = 1
}

func listOrder(pg *Graph) {
	for _, i := range pila {
		fmt.Print(pg.nodes[i].label, " ")
	}
	fmt.Println()
}

func findLeader(g *Graph) {
	for _, u := range g.nodes {
		if u.previo == -1 {
			fmt.Println(" Leader ", u.label)
		}
	}
}

func listTree(pg *Graph) {
	for _, i := range arbol {
		fmt.Print(pg.nodes[i].label, " ")
		nod := pg.nodes[i]
		for nod.previo != -1 {
			j := nod.previo
			fmt.Print(pg.nodes[j].label, " ")
			nod = pg.nodes[j]
		}
		fmt.Println()
	}
}

func readFile(name string, start time.Time) (lisnod []twoInt) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var anod twoInt
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
	return
}

func main() {
	start := time.Now()
	name := "./data/SCC.txt"
	if len(os.Args) > 1 {
		name = "./data/" + os.Args[1]
	}

	lisnod := readFile(name, start)
	elapsed := time.Since(start)
	fmt.Printf("Reading took %s\n", elapsed)

	nr := getMax(lisnod)
	fmt.Printf("Input %10d  Edges %10d Nodes\n", len(lisnod), nr)

	pg := new(Graph)
	pg.creatGraph(lisnod, false, start)
	elapsed = time.Since(start)
	fmt.Printf("Created %10d Nodes  %s\n", len(pg.nodes), elapsed)

	pg.display()

	pg.dfs(false)
	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)
	elapsed = time.Since(start)

	listOrder(pg)

	pgr := new(Graph)
	pgr.creatGraph(lisnod, true, start)
	elapsed = time.Since(start)
	fmt.Printf("Created %10d Nodes  %s\n", len(pg.nodes), elapsed)

	//pgr.display()

	pgr.dfs(true)
	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)
	elapsed = time.Since(start)
	//    findLeader(pgr)
	listTree(pgr)
}
