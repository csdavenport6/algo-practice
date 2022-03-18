package main

import (
	"algopractice/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAXV = 100 /* max number of verticies */

type Edgenode struct {
	y      int       /* adjacency info */
	weight int       /* edge weight, if any */
	next   *Edgenode /* next edge in list */
}

type Graph struct {
	edges     [MAXV + 1]*Edgenode /* adjacency info */
	degree    [MAXV + 1]int       /* out degree of each vertex */
	nvertices int
	nedges    int
	directed  bool
}

func main() {
	var G Graph
	ReadGraph(&G, false)
	PrintGraph(&G)
}

func InitializeGraph(g *Graph, directed bool) {
	var i int /* counter */

	g.nvertices = 0
	g.nedges = 0
	g.directed = directed

	for i = 1; i <= MAXV; i++ {
		g.degree[i] = 0
	}

	for i = 1; i <= MAXV; i++ {
		g.edges[i] = nil
	}
}

func ReadGraph(g *Graph, directed bool) {
	InitializeGraph(g, directed)

	fp := "graph_input.txt"
	f, err := os.Open(fp)
	helpers.Check(err)

	defer f.Close()

	r := bufio.NewScanner(f)
	r.Scan()

	fl := strings.Split(r.Text(), " ")
	n, err := strconv.Atoi(fl[0])
	helpers.Check(err)
	g.nvertices = n

	for r.Scan() {
		line := strings.Split(r.Text(), " ")

		x, err := strconv.Atoi(line[0])
		helpers.Check(err)
		y, err := strconv.Atoi(line[1])
		helpers.Check(err)

		InsertEdge(g, x, y, directed)
	}
}

func InsertEdge(g *Graph, x int, y int, directed bool) {
	var p Edgenode

	p.weight = 0
	p.y = y
	p.next = g.edges[x]

	g.edges[x] = &p

	g.degree[x]++

	if !directed {
		InsertEdge(g, y, x, true)
	} else {
		g.nedges++
	}
}

func PrintGraph(g *Graph) {
	for i := 1; i <= g.nvertices; i++ {
		fmt.Printf("%d: ", i)
		p := g.edges[i]
		for p != nil {
			fmt.Printf(" %d", p.y)
			p = p.next
		}
		fmt.Printf("\n")
	}
}
