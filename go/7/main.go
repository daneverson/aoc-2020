package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatalf("failed reading input file: %v", err)
	}

	graph, err := NewGraphFromAOC(strings.Split(string(contents), "\n"))
	if err != nil {
		log.Fatalf("failed to initialize graph: %v", err)
	}

	fmt.Printf("%s\n", graph.Print())

	startNode, err := graph.GetNodeByName("shiny gold")
	if err != nil {
		log.Fatalf("unable to find starting node: %v", err)
	}
	fmt.Printf("%d\n", findAllContainers(graph, startNode, 0, []*Edge{}))
	fmt.Printf("%d\n", countTotalBags(graph, startNode, []*Edge{}))
}

func countTotalBags(g *Graph, start *Node, q []*Edge) int {
	start.visited = true
	total := 0

	for _, e := range g.GetContained(start) {
		total = total + e.weight + (e.weight * countTotalBags(g, e.b, q))
	}

	return total
}

func findAllContainers(g *Graph, start *Node, total int, q []*Edge) int {
	start.visited = true
	q = append(q, g.GetContainers(start)...)

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		if !e.a.visited {
			total = findAllContainers(g, e.a, total+1, q)
		}
	}

	return total
}
