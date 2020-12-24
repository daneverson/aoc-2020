package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Graph is a collection of Nodes and Edges that is capable of some typical graph operations.
type Graph struct {
	nodes []*Node
	edges []*Edge
}

// Node is a graph vertex.
type Node struct {
	name    string
	visited bool
}

// Edge connects two Nodes and is directional from Node a to Node b.
type Edge struct {
	a, b   *Node
	weight int
}

// GetNodeByName returns the Node in this graph with the given name, or an error if not found.
func (g *Graph) GetNodeByName(name string) (*Node, error) {
	for _, n := range g.nodes {
		if n.name == name {
			return n, nil
		}
	}
	return nil, fmt.Errorf("Node with name '%s' was not found in this Graph", name)
}

// AddNode adds a Node to the Graph with the given name.
func (g *Graph) AddNode(name string) *Node {
	n := &Node{
		name:    name,
		visited: false,
	}
	g.nodes = append(g.nodes, n)
	return n
}

// AddEdge adds an Edge to the Graph directed from a to b, with a weight of w.
func (g *Graph) AddEdge(a, b *Node, w int) *Edge {
	e := &Edge{
		a:      a,
		b:      b,
		weight: w,
	}
	g.edges = append(g.edges, e)
	return e
}

// GetContained returns a slice of edges leading from the given Node to adjacent Nodes.
func (g *Graph) GetContained(n *Node) []*Edge {
	adjacent := []*Edge{}
	for _, e := range g.edges {
		if e.a.name == n.name {
			adjacent = append(adjacent, e)
		}
	}
	return adjacent
}

// GetContainers returns a slice of edges that lead from adjacent Nodes to the given Node.
func (g *Graph) GetContainers(n *Node) []*Edge {
	adjacent := []*Edge{}
	for _, e := range g.edges {
		if e.b.name == n.name {
			adjacent = append(adjacent, e)
		}
	}
	return adjacent
}

// Print returns a string representation of the Graph suitable for printing to stdout.
func (g *Graph) Print() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "digraph G {\n")
	for _, n := range g.nodes {
		for _, e := range g.GetContained(n) {
			fmt.Fprintf(b, "  \"%s\" -> \"%s\" [label=%d];\n", n.name, e.b.name, e.weight)
		}
	}
	fmt.Fprintf(b, "}\n")
	return b.String()
}

// NewGraphFromAOC takes a slice of rule strings and parses them into a set of Nodes and Edges.
// The expected input is as described in the advent of code 2020 problem for day 7, e.g.:
// > "clear gold bags contain 5 vibrant gray bags, 5 wavy white bags."
func NewGraphFromAOC(rules []string) (*Graph, error) {
	g := Graph{
		nodes: []*Node{},
		edges: []*Edge{},
	}

	for _, rule := range rules {
		containRule := strings.Split(rule, " bags contain ")

		// The container Node - either get the existing one if we've already seen a Node by this
		// name, or create it if not.
		containerColor := containRule[0]
		containerNode, err := g.GetNodeByName(containerColor)
		if err != nil {
			containerNode = g.AddNode(containerColor)
		}

		// A container Node may have zero or more contained Nodes.
		containedNodes := strings.Split(containRule[1], ", ")
		for _, containedNode := range containedNodes {
			// Get the color and number of contained bags for entry
			s := strings.Split(containedNode, " ")
			if s[0] == "no" {
				// this would be if the bag description includes "contain no other bags"
				continue
			}
			containedAmount, err := strconv.ParseUint(s[0], 10, 64)
			if err != nil {
				return nil, fmt.Errorf(
					"Input was invalid - '%s' found where an integer was expected", s[0])
			}
			containedColor := strings.Join(s[1:3], " ")

			// Retrieve existing or create a new Node for the contained element, then add a weighted
			// edge connecting this Node to its parent.
			containedNode, err := g.GetNodeByName(containedColor)
			if err != nil {
				containedNode = g.AddNode(containedColor)
			}
			g.AddEdge(containerNode, containedNode, int(containedAmount))
		}
	}

	return &g, nil
}
