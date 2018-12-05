package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	label   int
	edges   []int
	visited bool
}

func (n Node) String() string {
	return fmt.Sprintf("id:\t%d\nodos: %v\n\n", n.label, n.edges)
}

func (n *Node) AddEdge(i int) {
	n.edges = append(n.edges, i)
}

type Group struct {
	label   int
	members map[int]bool
}

func (g *Group) AddMember(i int) {
	g.members[i] = true
}

var mapF = make(map[int]*Node)
var mapB = make(map[int]*Node)
var keyMap = make(map[int]int)
var ordenMap = make(map[int]int)
var mainMap = make(map[int]*Group)

var s *Group

var t int

func main() {

	var name string

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ReadFile(name)

	DFSLoop(mapB, keyMap, 1)
	DFSLoop(mapF, ordenMap, 2)

	for _, v := range mainMap {
		for w := range v.members {
			_, ok := v.members[w*-1]

			if ok {
				fmt.Println("No tiene solucion")
				os.Exit(1)
			}
		}
	}

	fmt.Println("Tiene solucion")

}

func ReadFile(filename string) {

	k := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		numOfRows, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatalf("No se pudo convertir el numero: %v\n", err)
		}

		for i := 1; i <= numOfRows; i++ {

			ni := i * -1
			vf := &Node{i, []int{}, false}
			mapF[i] = vf

			nvf := &Node{ni, []int{}, false}
			mapF[ni] = nvf

			vb := &Node{i, []int{}, false}
			mapB[i] = vb

			nvb := &Node{ni, []int{}, false}
			mapB[ni] = nvb

			keyMap[k] = i
			k++
			keyMap[k] = ni
			k++
		}
	}

	for scanner.Scan() {

		thisLine := strings.Fields(scanner.Text())

		sat1, err := strconv.Atoi(thisLine[0])
		sat2, err := strconv.Atoi(thisLine[1])

		if err != nil {
			fmt.Print("No se puede conventir el numero: %v\n", err)
			return
		}

		nsat1V, ok := mapF[sat1*-1]

		if !ok {
			log.Fatal("No puede ser encontrado: ", sat1*-1)
		}
		nsat2V, ok := mapF[sat2*-1]

		if !ok {
			log.Fatal("No puede ser encontrado: ", sat1*-1)
		}

		nsat1V.AddEdge(sat2)
		nsat2V.AddEdge(sat1)

		sat1V, ok := mapB[sat1]

		if !ok {
			log.Fatal("No puede ser encontrado: ", sat1)
		}

		sat2V, ok := mapB[sat2]

		if !ok {
			log.Fatal("No puede ser encontrado: ", sat2)
		}

		sat1V.AddEdge(sat2 * -1)
		sat2V.AddEdge(sat1 * -1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func DFSLoop(graph map[int]*Node, orderer map[int]int, pass int) {

	for i := len(orderer); i > 0; i-- {

		w, ok := graph[orderer[i]]

		if ok {

			if !w.visited {

				if pass == 2 {
					s = &Group{w.label, make(map[int]bool)}
					mainMap[s.label] = s
				}

				DFS(graph, w, pass)
			}

		}
	}
}

func DFS(graph map[int]*Node, i *Node, pass int) {

	i.visited = true

	if pass == 2 {
		s.AddMember(i.label)
	}

	for _, v := range i.edges {

		Vertice := graph[v]

		if !Vertice.visited {
			DFS(graph, Vertice, pass)
		}
	}

	if pass == 1 {
		t++
		ordenMap[t] = i.label
	}

}
