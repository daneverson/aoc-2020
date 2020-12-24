package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type set map[string]struct{}

func main() {
	contents, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatalf("failed reading input file: %v", err)
	}

	groups := strings.Split(string(contents), "\n\n")

	totalYes := 0
	for i, g := range groups {
		groupYes := makeSet(g)
		fmt.Printf("group %d: %d yes answers\n", i, len(groupYes))
		totalYes = totalYes + len(groupYes)
	}

	fmt.Printf("%d\n", totalYes)
}

func makeSet(group string) map[string]struct{} {
	answerLists := strings.Split(group, "\n")
	sets := make([]set, len(answerLists))
	var returnSet set

	for i, answerList := range answerLists {
		sets[i] = make(set)
		for _, answer := range answerList {
			sets[i][string(answer)] = struct{}{}
		}
		if i == 0 {
			returnSet = sets[i]
		} else {
			returnSet = setIntersect(returnSet, sets[i])
		}
	}

	return returnSet
}

func setIntersect(a, b set) set {
	intersection := set{}
	for el := range a {
		if _, ok := b[el]; ok {
			intersection[el] = struct{}{}
		}
	}
	return intersection
}
