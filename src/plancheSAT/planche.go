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
	Id       int
	Edges    []int
	Explored bool
}

func (v Vertice) String() string {
	return fmt.Sprintf("id:\t%d\nodos: %v\n\n", v.Id, v.Edges)
}

func (v *Vertice) AddEdge(i int) {
	v.Edges = append(v.Edges, i)
}

type Boss struct {
	Id       int
	Miembros map[int]bool
}

func (v Boss) String() string {
	return fmt.Sprintf("id:\t%d\n miembros: %v\n\n", v.Id, v.Miembros)
}

func (v *Boss) AddMember(i int) {
	v.Miembros[i] = true
}

var VerticeMap_f = make(map[int]*Vertice)
var VerticeMap_b = make(map[int]*Vertice)
var Atras_Key = make(map[int]int)
var m_Orden_Map = make(map[int]int)
var Main_map = make(map[int]*Boss)

var s *Boss

var t int

func main() {

	var name string

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	renombrarFile(name)

	bucleDFS(VerticeMap_b, Atras_Key, 1)
	bucleDFS(VerticeMap_f, m_Orden_Map, 2)

	for _, v := range Main_map {
		for w := range v.Miembros {
			_, ok := v.Miembros[w*-1]

			if ok {
				fmt.Println("La satisfacibilidad es 0")
				os.Exit(1)
			}
		}
	}

	fmt.Println("La satisfacibilidad es 1")

}

func renombrarFile(filename string) {

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
			log.Fatalf("no se pudo comvertir el número: %v\n", err)
		}

		for i := 1; i <= numOfRows; i++ {

			ni := i * -1
			vf := &Vertice{i, []int{}, false}
			VerticeMap_f[i] = vf

			nvf := &Vertice{ni, []int{}, false}
			VerticeMap_f[ni] = nvf

			vb := &Vertice{i, []int{}, false}
			VerticeMap_b[i] = vb

			nvb := &Vertice{ni, []int{}, false}
			VerticeMap_b[ni] = nvb

			Atras_Key[k] = i
			k++
			Atras_Key[k] = ni
			k++
		}
	}

	for scanner.Scan() {

		thisLine := strings.Fields(scanner.Text())

		sat1, err := strconv.Atoi(thisLine[0])
		sat2, err := strconv.Atoi(thisLine[1])

		if err != nil {
			fmt.Print("No se puede conventir el número(Can't convert the number): %v\n", err)
			return
		}

		nsat1V, ok := VerticeMap_f[sat1*-1]

		if !ok {
			log.Fatal("Couldn't Find it (No puede ser encontrado): ", sat1*-1)
		}
		nsat2V, ok := VerticeMap_f[sat2*-1]

		if !ok {
			log.Fatal("Couldn't Find it(No puede ser encontrado): ", sat1*-1)
		}

		nsat1V.AddEdge(sat2)
		nsat2V.AddEdge(sat1)

		sat1V, ok := VerticeMap_b[sat1]

		if !ok {
			log.Fatal("No puede ser encontrado(Couldn't Find it): ", sat1)
		}

		sat2V, ok := VerticeMap_b[sat2]

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

func bucleDFS(graph map[int]*Vertice, orderer map[int]int, pass int) {

	for i := len(orderer); i > 0; i-- {

		w, ok := graph[orderer[i]]

		if ok {

			if !w.Explored {

				if pass == 2 {
					s = &Boss{w.Id, make(map[int]bool)}
					Main_map[s.Id] = s
				}

				DFS(graph, w, pass)
			}

		}
	}
}

func DFS(graph map[int]*Vertice, i *Vertice, pass int) {

	i.Explored = true

	if pass == 2 {
		s.AddMember(i.Id)
	}

	for _, v := range i.Edges {

		Vertice := graph[v]

		if !Vertice.Explored {
			DFS(graph, Vertice, pass)
		}
	}

	if pass == 1 {
		t++
		m_Orden_Map[t] = i.Id
	}

}
