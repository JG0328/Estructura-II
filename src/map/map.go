package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	label string
	color string
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.label)
}

func (g *Graph) getNode(st string) *Node {
	for _, pnd := range g.nodes {
		if pnd.String() == st {
			return pnd
		}
	}
	return nil
}

func (g *Graph) display() {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].label + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].label + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}

func seekCountry(ct string, ls []string) (b bool) {
	b = false
	for _, st := range ls {
		if st == ct {
			b = true
			break
		}
	}
	return b
}

func getLisCountries(lsls [][]string) []string {
	var lis []string
	for _, stlis := range lsls {
		for _, st := range stlis {
			if !seekCountry(st, lis) {
				lis = append(lis, st)
			}
		}
	}
	return lis
}

func (pg *Graph) creatGraph(lst []string, lsls [][]string) {
	for _, st := range lst {
		nd := new(Node)
		nd.label = st
		pg.AddNode(nd)
	}
	for _, lis := range lsls {
		pini := pg.FindNode(lis[0])

		for _, st := range lis[1:] {
			nd := pg.FindNode(st)
			pg.AddEdge(pini, nd)
		}
	}
}

// Función que busca un nodo
func (g *Graph) FindNode(label string) *Node {
	for i := 0; i < len(g.nodes); i++ {
		if label == g.nodes[i].label {
			return g.nodes[i]
		}
	}
	return nil
}

// Se agrega a la lista, sin repetir
func AppendNoRepeat(list []string, str string) string {
	for i := 0; i < len(list); i++ {
		if str == list[i] {
			return ""
		}
	}
	return str
}

// ¿Se puede colorear este nodo, a partir de una lista de colores prohibidos?
func (node *Node) CanColor(noColors []string) bool {
	for i := 0; i < len(noColors); i++ {
		if noColors[i] == node.label {
			return false
		}
	}
	return true
}

// Función que se encarga de colorear el mapa
func (g *Graph) ColorMap() {
	for i := 0; i < len(g.nodes); i++ {
		var noColors []string
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			noColors = append(noColors, AppendNoRepeat(noColors, near[j].color))
		}
		// Colores (países)
		for j := 0; j < len(g.nodes); j++ {
			if g.nodes[j].CanColor(noColors) {
				g.nodes[i].color = g.nodes[j].label
				break
			}
		}
	}
	var lislis [][]string
	for i := 0; i < len(g.nodes); i++ {
		for j := 0; j < len(g.nodes); j++ {
			if g.nodes[i].color == g.nodes[j].label {
				if j >= len(lislis) || len(lislis) == 0 {
					var lis []string
					lislis = append(lislis, lis)
				}
				lislis[j] = append(lislis[j], g.nodes[i].label)
				break
			}
		}
	}
	// Imprimir
	for i := 0; i < len(lislis); i++ {
		for j := 0; j < len(g.nodes); j++ {
			for k := 0; k < len(lislis[i]); k++ {
				if g.nodes[j].label == lislis[i][k] && k < len(lislis[i])-1 {
					fmt.Printf("(%s,%s), ", g.nodes[j].label, g.nodes[j].color)
				} else if g.nodes[j].label == lislis[i][k] && k == len(lislis[i])-1 {
					fmt.Printf("(%s,%s).", g.nodes[j].label, g.nodes[j].color)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Printf("Se utilizaron: %d colores.\n\n", len(lislis))
}
func main() {
	file, err := os.Open("eu_map.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var g Graph
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var lislis [][]string
	for scanner.Scan() {
		lineStr := scanner.Text()
		list := strings.Split(lineStr, " ")
		lislis = append(lislis, list)
	}
	fmt.Println()
	lisc := getLisCountries(lislis)
	(&g).creatGraph(lisc, lislis)
	(&g).ColorMap()
}
