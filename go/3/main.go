package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	TREE   = '#'
	NOTREE = '.'
)

func main() {
	contents, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatalf("failed reading input file: %v", err)
	}
	lines := strings.Split(string(contents), "\n")

	var slopesToCheck = []struct {
		x, y int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	results := make(chan int, len(slopesToCheck))

	for _, slope := range slopesToCheck {
		trees := treesEncountered(slope.x, slope.y, lines)
		fmt.Printf("%d trees encountered\n", trees)
		results <- trees
	}
	close(results)

	acc := 1
	for t := range results {
		acc = acc * t
	}
	fmt.Println(acc)
}

func treesEncountered(x, y int, lines []string) int {
	numTrees := 0
	wrapWidth := len(lines[0])
	for i := y; i < len(lines); i = i + y {
		xValue := (i / y) * x % wrapWidth
		if lines[i][xValue] == TREE {
			numTrees++
		}
	}
	return numTrees
}
