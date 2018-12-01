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

// Código facilitado por el profesor
type Node struct {
	label     int
	visited   bool
	neighbors []*Node
}

type Graph struct {
	nodes []*Node
}

/*
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
*/

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

	for i := 0; i < len(g.nodes); i++ {
		if g.nodes[i].label == label {
			return g.nodes[i]
		}
	}

	return nil
}

func CreateNode(label int) *Node {
	n := new(Node)
	n.label = label
	return n
}

func (g *Graph) AddEdge(nini *Node, nfin *Node) {
	nini.neighbors = append(nini.neighbors, nfin)
}

func CreateGraph(bytesRead []byte) *Graph {
	nodes := strings.Fields(string(bytesRead))

	start := time.Now()

	g := new(Graph)

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
		} else {
			nini = g.GetNode(labelIni)
		}
		if g.GetNode(labelFin) == nil {
			nfin = CreateNode(labelFin)
		} else {
			nfin = g.GetNode(labelFin)
		}

		g.AddEdge(nini, nfin)
	}

	var elapsed time.Duration
	elapsed = time.Since(start)

	fmt.Printf("Creation took %s\n", elapsed)

	return g
}

func (g *Graph) AddNode(node *Node) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) Display() {
	/*
		s := ""
		for i := 0; i < len(g.nodes); i++ {
			s += fmt.Sprintf("%8d --> ", g.nodes[i].label)
			near := g.nodes[i]
			for j := 0; j < len(near.neighbors); j++ {
				s += fmt.Sprintf("%8d ", near.neighbors[j].label)
			}
			//        s += "\n"
			fmt.Println(s)
			s = ""
		}
	*/
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

	gr := CreateGraph(bytesRead)

	if gr == nil {
		return
	}

	gr.Display()

	/*
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
	*/
}
