package main

import (
	"fmt"
	"os"
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

//

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

var adj map[int][]int
var adjInv map[int][]int
var visited map[int]bool
var visitedInv map[int]bool

var scc map[int]int
var counter int

func main() {
	var n string
	if len(os.Args) > 1 {
		n = os.Args[1]
	}

	s := NewStack()

	// Leo el archivo y obtengo los datos
	bytes := GetFile(n)

	words := strings.Fields(string(bytes))

	var a []int
	var b []int

	for i := 0; i < len(words); i += 2 {
		toA, err := strconv.Atoi(words[i])
		toB, err2 := strconv.Atoi(words[i+1])

		if err != nil || err2 != nil {
			fmt.Println("ERROR GETTING DATA")
			return
		}

		a = append(a, toA)
		b = append(b, toB)
	}

	adj = make(map[int][]int)
	adjInv = make(map[int][]int)
	visited = make(map[int]bool)
	visitedInv = make(map[int]bool)
	scc = make(map[int]int)

	counter = 1

	cantVars := make(map[int]bool)

	for i := 0; i < len(words); i++ {
		sti, err := strconv.Atoi(words[i])

		if err != nil {
			fmt.Println("ERROR GETTING DATA")
			return
		}

		if cantVars[sti] == false && sti > 0 {
			cantVars[sti] = true
		}
	}

	fmt.Println("CANTIDAD DE VARIABLES -> ", len(cantVars))

	SAT(len(cantVars), len(a), a, b, s)

	// fmt.Print(a[0], b[0])
}

func SAT(n int, m int, a []int, b []int, s *Stack) {
	for i := 0; i < m; i++ {
		if a[i] > 0 && b[i] > 0 {
			addEdges(a[i]+n, b[i])
			addEdgesInverse(a[i]+n, b[i])
			addEdges(b[i]+n, a[i])
			addEdgesInverse(b[i]+n, a[i])
		} else if a[i] > 0 && b[i] < 0 {
			addEdges(a[i]+n, n-b[i])
			addEdgesInverse(a[i]+n, n-b[i])
			addEdges(-b[i], a[i])
			addEdgesInverse(-b[i], a[i])
		} else if a[i] < 0 && b[i] > 0 {
			addEdges(-a[i], b[i])
			addEdgesInverse(-a[i], b[i])
			addEdges(b[i]+n, n-a[i])
			addEdgesInverse(b[i]+n, n-a[i])
		} else {
			addEdges(-a[i], n-b[i])
			addEdgesInverse(-a[i], n-b[i])
			addEdges(-b[i], n-a[i])
			addEdgesInverse(-b[i], n-a[i])
		}
	}

	for i := 1; i <= 2*n; i++ {
		if !visited[i] {
			dfsFirst(i, s)
		}
	}

	for {
		if s.Len() == 0 {
			break
		}
		n = (s.Peek()).(int)
		s.Pop()

		if !visitedInv[n] {
			dfsSecond(n, s)
			counter++
		}
	}

	for i := 1; i <= n; i++ {
		// fmt.Println(scc[i], " y ", scc[i+n])
		if scc[i] == scc[i+n] {
			fmt.Println("No tiene solucion")
			return
		}
	}

	// fmt.Println("ADJ MAP -> ", adj)

	fmt.Println("Tiene solucion")
	return
}

func dfsFirst(u int, s *Stack) {
	if visited[u] {
		return
	}

	visited[u] = true

	for i := 0; i < len(adj[u]); i++ {
		dfsFirst(adj[u][i], s)
	}

	s.Push(u)
}

func dfsSecond(u int, s *Stack) {
	if visitedInv[u] {
		return
	}

	visitedInv[u] = true

	for i := 0; i < len(adjInv[u]); i++ {
		dfsSecond(adjInv[u][i], s)
	}

	scc[u] = counter
}

func addEdges(a int, b int) {
	adj[a] = append(adj[a], b)
}

func addEdgesInverse(a int, b int) {
	adjInv[b] = append(adjInv[b], a)
}
