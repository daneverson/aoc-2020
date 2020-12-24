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
		log.Fatalf("could not read input file: %v", err)
	}
	tokens := strings.Split(string(contents), "\n")
	for i, t1 := range tokens {
		for j, t2 := range tokens {
			for k, t3 := range tokens {
				v1, err := strconv.ParseInt(t1, 10, 64)
				if err != nil {
					log.Fatalf("encountered a non-integer on line %d: %s", i, t1)
				}
				v2, err := strconv.ParseInt(t2, 10, 64)
				if err != nil {
					log.Fatalf("encountered a non-integer on line %d: %s", j, t2)
				}
				v3, err := strconv.ParseInt(t3, 10, 64)
				if err != nil {
					log.Fatalf("encountered a non-integer on line %d: %s", k, t2)
				}
				if v1+v2+v3 == 2020 {
					fmt.Printf("Found result: %d\n", v1*v2*v3)
					return
				}
			}
		}
	}
}
