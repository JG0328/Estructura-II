package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vertice struct {
	id      int
	edges   []int
	visited bool
}

func (v *Vertice) AddEdge(i int) {
	v.edges = append(v.edges, i)
}

type Boss struct {
	id      int
	members map[int]bool
}

func (v *Boss) AddMember(i int) {
	v.members[i] = true
}

func readFile(filename string) {
	k := 0

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		rows, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatalf("no se pudo convertir el número: %v\n", err)
		}

		for i := 1; i <= rows; i++ {

			ni := i * -1
			vf := &Vertice{i, []int{}, false}
			mapF[i] = vf

			nvf := &Vertice{ni, []int{}, false}
			mapF[ni] = nvf

			vb := &Vertice{i, []int{}, false}
			mapB[i] = vb

			nvb := &Vertice{ni, []int{}, false}
			mapB[ni] = nvb

			keyMap[k] = i
			k++
			keyMap[k] = ni
			k++
		}
	}

	for scanner.Scan() {

		line := strings.Fields(scanner.Text())

		sat1, err := strconv.Atoi(line[0])
		sat2, err := strconv.Atoi(line[1])

		if err != nil {
			fmt.Print("No se puede conventir el número(Can't convert the number): %v\n", err)
			return
		}

		nsat1V, ok := mapF[sat1*-1]

		if !ok {
			log.Fatal("Couldn't Find it (No puede ser encontrado): ", sat1*-1)
		}
		nsat2V, ok := mapF[sat2*-1]

		if !ok {
			log.Fatal("Couldn't Find it(No puede ser encontrado): ", sat1*-1)
		}

		nsat1V.AddEdge(sat2)
		nsat2V.AddEdge(sat1)

		sat1V, ok := mapB[sat1]

		if !ok {
			log.Fatal("No puede ser encontrado(Couldn't Find it): ", sat1)
		}

		sat2V, ok := mapB[sat2]

		if !ok {
			log.Fatal("No puede ser encontrado(Couldn't Find it): ", sat2)
		}

		sat1V.AddEdge(sat2 * -1)
		sat2V.AddEdge(sat1 * -1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func cicloDFS(graph map[int]*Vertice, orderer map[int]int, pass int) {

	for i := len(orderer); i > 0; i-- {

		w, ok := graph[orderer[i]]

		if ok {
			if !w.visited {
				if pass == 2 {
					g = &Boss{w.id, make(map[int]bool)}
					mainMap[g.id] = g
				}
				dfs(graph, w, pass)
			}
		}
	}
}

func dfs(graph map[int]*Vertice, i *Vertice, pass int) {
	i.visited = true

	if pass == 2 {
		g.AddMember(i.id)
	}

	for _, v := range i.edges {

		node := graph[v]

		if !node.visited {
			dfs(graph, node, pass)
		}
	}

	if pass == 1 {
		t++
		orderMap[t] = i.id
	}
}

var mapF = make(map[int]*Vertice)
var mapB = make(map[int]*Vertice)
var keyMap = make(map[int]int)
var orderMap = make(map[int]int)
var mainMap = make(map[int]*Boss)
var g *Boss
var t int

func main() {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	readFile(name)
	cicloDFS(mapB, keyMap, 1)
	cicloDFS(mapF, orderMap, 2)

	for _, v := range mainMap {
		for w := range v.members {
			_, ok := v.members[w*-1]

			if ok {
				fmt.Println("La satisfacibilidad es 0")
				return
			}
		}
	}
	fmt.Println("La satisfacibilidad es 1")
}
