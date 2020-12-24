package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatalf("failed reading input file: %v", err)
	}

	split := strings.Split(string(contents), "\n")
	numList := make([]uint64, len(split))
	for i, n := range split {
		parsed, _ := strconv.ParseInt(n, 10, 64)
		numList[i] = uint64(parsed)
	}

	fmt.Println(problemOne(numList))
}

func problemOne(n []uint64) uint64 {
	return 0
}
