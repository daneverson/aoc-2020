package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type validatorFunc func(int64, int64, rune, string) bool

var (
	validator validatorFunc
)

func main() {
	contents, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatalf("failed reading input file: %v", err)
	}

	lines := strings.Split(string(contents), "\n")
	validPasswords := 0

	validator = tobogganValidation
	for _, l := range lines {
		min, max, requiredLetter, rest := parseReqs(l)
		if validator(min, max, requiredLetter, rest) {
			validPasswords++
		}
	}

	fmt.Printf("Found %d valid passwords\n", validPasswords)
}

func tobogganValidation(pos1, pos2 int64, requiredLetter rune, pass string) bool {
	return (rune(pass[pos1-1]) == requiredLetter) != (rune(pass[pos2-1]) == requiredLetter)
}

func sledRentalValidation(min, max int64, requiredLetter rune, pass string) bool {
	seen := int64(0)
	for _, c := range pass {
		if c == requiredLetter {
			seen++
			if seen > max {
				break
			}
		}
	}
	if seen <= max && seen >= min {
		return true
	}
	return false
}

func parseReqs(line string) (int64, int64, rune, string) {
	prefixAndPass := strings.Split(line, ":")
	splitPrefix := strings.Split(prefixAndPass[0], " ")
	minMax := strings.Split(splitPrefix[0], "-")
	min, _ := strconv.ParseInt(minMax[0], 10, 64)
	max, _ := strconv.ParseInt(minMax[1], 10, 64)
	return min, max, rune(splitPrefix[1][0]), strings.Trim(prefixAndPass[1], " ")
}
